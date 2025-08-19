package values

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_UIntUnwrapTo(t *testing.T) {
	expected := uint64(100)
	v := NewUint64(expected)

	var got int
	err := v.UnwrapTo(&got)
	require.NoError(t, err)

	assert.Equal(t, expected, uint64(got))

	var gotUint64 uint64
	err = v.UnwrapTo(&gotUint64)
	require.NoError(t, err)

	assert.Equal(t, expected, gotUint64)

	var varAny any
	err = v.UnwrapTo(&varAny)
	require.NoError(t, err)
	assert.Equal(t, expected, varAny)

	in := (*Uint64)(nil)
	_, err = in.Unwrap()
	assert.ErrorContains(t, err, "cannot unwrap nil")

	var i uint64
	err = in.UnwrapTo(&i)
	assert.ErrorContains(t, err, "cannot unwrap nil")
}

func Test_UintUnwrapping(t *testing.T) {
	t.Run("uint64 -> uint32", func(st *testing.T) {
		expected := uint64(100)
		v := NewUint64(expected)
		got := uint32(0)
		err := v.UnwrapTo(&got)
		require.NoError(t, err)
		assert.Equal(t, uint32(expected), got)
	})

	t.Run("uint64 -> uint32; overflow", func(st *testing.T) {
		expected := uint64(math.MaxUint64)
		v := NewUint64(expected)
		got := uint32(0)
		err := v.UnwrapTo(&got)
		assert.NotNil(t, err)
	})

	t.Run("uint64 -> int32", func(st *testing.T) {
		expected := uint64(100)
		v := NewUint64(expected)
		got := int32(0)
		err := v.UnwrapTo(&got)
		require.NoError(t, err)
		assert.Equal(t, int32(expected), got)
	})

	t.Run("uint64 -> int32; overflow", func(st *testing.T) {
		expected := uint64(math.MaxInt64)
		v := NewUint64(expected)
		got := uint32(0)
		err := v.UnwrapTo(&got)
		assert.NotNil(t, err)
	})

	t.Run("uint64 -> int64; overflow", func(st *testing.T) {
		expected := uint64(math.MaxUint64)
		v := NewUint64(expected)
		got := int64(0)
		err := v.UnwrapTo(&got)
		assert.NotNil(t, err)
	})
}
