package values

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Float64UnwrapTo(t *testing.T) {
	expected := 1.1
	v := NewFloat64(expected)

	var got float64
	err := v.UnwrapTo(&got)
	require.NoError(t, err)

	assert.Equal(t, expected, got)

	gotn := (*float64)(nil)
	err = v.UnwrapTo(gotn)
	assert.ErrorContains(t, err, "cannot unwrap to nil pointer")

	var varAny any
	err = v.UnwrapTo(&varAny)
	require.NoError(t, err)
	assert.Equal(t, expected, varAny)

	fn := (*Float64)(nil)
	_, err = fn.Unwrap()
	assert.ErrorContains(t, err, "cannot unwrap nil")

	var f float64
	err = fn.UnwrapTo(&f)
	assert.ErrorContains(t, err, "cannot unwrap nil")

	// handle alias
	type myFloat64 float64
	var mf myFloat64
	err = v.UnwrapTo(&mf)
	require.NoError(t, err)
	assert.Equal(t, myFloat64(expected), mf)
}

func Test_Float64UnwrapToFloat32(t *testing.T) {
	t.Run("withing range of float 32 allowing precision loss", func(t *testing.T) {
		expected := 0.123456789012345
		v := NewFloat64(expected)

		var got float32
		err := v.UnwrapTo(&got)
		require.NoError(t, err)

		assert.Equal(t, float32(expected), got)
		// ensure the number we're using is captures precision loss
		assert.NotEqual(t, expected, float64(got))
	})

	t.Run("too large for a float 32", func(t *testing.T) {
		expected := 3.5e38
		v := NewFloat64(expected)

		var got float32
		err := v.UnwrapTo(&got)
		require.Error(t, err)
	})

	t.Run("too small for a float 32", func(t *testing.T) {
		expected := 1e-50
		v := NewFloat64(expected)

		var got float32
		err := v.UnwrapTo(&got)
		require.Error(t, err)
	})

	t.Run("alias type", func(t *testing.T) {
		type myfloat32 float32
		expected := myfloat32(0.123456789012345)
		v := NewFloat64(float64(expected))
		var got myfloat32
		err := v.UnwrapTo(&got)
		require.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

func Test_Float64UnwrapToDecimal(t *testing.T) {
	t.Run("unwraps to decimal", func(t *testing.T) {
		expected := decimal.NewFromFloat(1.1)
		v := NewFloat64(1.1)

		var got decimal.Decimal
		err := v.UnwrapTo(&got)
		require.NoError(t, err)
		require.Equal(t, expected, got)

		assert.Equal(t, expected, got)

		gotn := (*decimal.Decimal)(nil)
		err = v.UnwrapTo(gotn)
		assert.ErrorContains(t, err, "cannot unwrap to nil pointer")
	})

	t.Run("unwraps to decimal alias", func(t *testing.T) {
		type myDecimal decimal.Decimal
		expected := myDecimal(decimal.NewFromFloat(1.1))
		v := NewFloat64(1.1)

		var got myDecimal
		err := v.UnwrapTo(&got)
		require.NoError(t, err)
		require.Equal(t, expected, got)

		assert.Equal(t, expected, got)

		gotn := (*myDecimal)(nil)
		err = v.UnwrapTo(gotn)
		assert.ErrorContains(t, err, "cannot unwrap to nil pointer")
	})
}

// Test_Float64 tests that Float64 values can converted to and from protobuf representations.
func Test_Float64(t *testing.T) {
	testCases := []struct {
		name string
		f    float64
	}{
		{
			name: "positive",
			f:    1.1,
		},
		{
			name: "0",
			f:    0,
		},
		{
			name: "negative",
			f:    -1.1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(st *testing.T) {
			v := NewFloat64(tc.f)

			vp := Proto(v)
			got, err := FromProto(vp)
			assert.NoError(t, err)
			assert.Equal(t, tc.f, got.(*Float64).Underlying)
		})
	}
}
