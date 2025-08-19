package values

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
)

type List struct {
	Underlying []Value
}

func NewList(l []any) (*List, error) {
	lv := []Value{}
	for _, v := range l {
		ev, err := Wrap(v)
		if err != nil {
			return nil, err
		}

		lv = append(lv, ev)
	}
	return &List{Underlying: lv}, nil
}

func (l *List) proto() *pb.Value {
	v := []*pb.Value{}
	for _, e := range l.Underlying {
		v = append(v, Proto(e))
	}
	return pb.NewListValue(v)
}

func (l *List) Unwrap() (any, error) {
	nl := []any{}
	return nl, l.UnwrapTo(&nl)
}

func (l *List) copy() Value {
	return l.CopyList()
}

func (l *List) CopyList() *List {
	if l == nil {
		return nil
	}

	dest := []Value{}
	for _, el := range l.Underlying {
		dest = append(dest, Copy(el))
	}
	return &List{Underlying: dest}
}

func (l *List) UnwrapTo(to any) error {
	if l == nil {
		return errors.New("cannot unwrap nil values.List")
	}

	val := reflect.ValueOf(to)
	if val.Kind() != reflect.Pointer {
		return fmt.Errorf("cannot unwrap to non-pointer type %T", to)
	}

	if val.IsNil() {
		return fmt.Errorf("cannot unwrap to nil pointer: %+v", to)
	}

	ptrVal := reflect.Indirect(val)
	switch ptrVal.Kind() {
	case reflect.Slice:
		newList := reflect.MakeSlice(ptrVal.Type(), len(l.Underlying), len(l.Underlying))
		return l.unwrapToSliceOrArray(newList, val)
	case reflect.Array:
		if ptrVal.Len() < len(l.Underlying) {
			return fmt.Errorf("too many elements to unwrap")
		} else if ptrVal.Len() > len(l.Underlying) {
			return fmt.Errorf("too few elements to unwrap")
		}
		arr := reflect.New(ptrVal.Type()).Elem()
		return l.unwrapToSliceOrArray(arr, val)
	default:
		dl := []any{}
		err := l.UnwrapTo(&dl)
		if err != nil {
			return err
		}

		if reflect.TypeOf(dl).AssignableTo(ptrVal.Type()) {
			ptrVal.Set(reflect.ValueOf(dl))
			return nil
		}

		return fmt.Errorf("cannot unwrap to type %T", to)
	}
}

func (l *List) unwrapToSliceOrArray(newList reflect.Value, val reflect.Value) error {
	for i, el := range l.Underlying {
		newElm := newList.Index(i)
		if newElm.Kind() == reflect.Pointer {
			newElm.Set(reflect.New(newElm.Type().Elem()))
		} else {
			newElm = newElm.Addr()
		}

		if el == nil {
			continue
		}
		if err := el.UnwrapTo(newElm.Interface()); err != nil {
			return err
		}
	}
	reflect.Indirect(val).Set(newList)
	return nil
}
