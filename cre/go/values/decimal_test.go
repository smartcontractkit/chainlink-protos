package values

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DecimalUnwrapTo(t *testing.T) {
	dv := decimal.NewFromFloat(1.00)
	tr := NewDecimal(dv)

	var dec decimal.Decimal
	err := tr.UnwrapTo(&dec)
	require.NoError(t, err)

	assert.Equal(t, dv, dec)

	var s string
	err = tr.UnwrapTo(&s)
	require.Error(t, err)

	decn := (*decimal.Decimal)(nil)
	err = tr.UnwrapTo(decn)
	assert.ErrorContains(t, err, "unwrap to nil pointer")

	var varAny any
	err = tr.UnwrapTo(&varAny)
	require.NoError(t, err)
	assert.Equal(t, dv, varAny)

	dn := (*Decimal)(nil)
	_, err = dn.Unwrap()
	assert.ErrorContains(t, err, "could not unwrap nil")

	dec = decimal.Decimal{}
	err = dn.UnwrapTo(&dec)
	assert.ErrorContains(t, err, "could not unwrap nil")
}

func Test_DecimalUnwrapToFloat64(t *testing.T) {
	t.Run("Unwraps within range", func(t *testing.T) {
		expected := 1.33342
		d := decimal.NewFromFloat(expected)
		wrapped := NewDecimal(d)
		var actual float64
		err := wrapped.UnwrapTo(&actual)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Unwraps with alias type", func(t *testing.T) {
		type myfloat float64
		expected := myfloat(1.33342)
		d := decimal.NewFromFloat(float64(expected))
		wrapped := NewDecimal(d)
		var actual myfloat
		err := wrapped.UnwrapTo(&actual)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Unwraps with loss of precision", func(t *testing.T) {
		precisionLoss := decimal.RequireFromString("1.23456789012345678901234567890")
		wrapped := NewDecimal(precisionLoss)
		var actual float64
		err := wrapped.UnwrapTo(&actual)
		require.NoError(t, err)
		assert.Equal(t, precisionLoss.InexactFloat64(), actual)
	})

	t.Run("Too small for float64", func(t *testing.T) {
		tooSmall := decimal.RequireFromString("1e-400") // smaller than ~5e-324
		wrapped := NewDecimal(tooSmall)
		var actual float64
		require.Error(t, wrapped.UnwrapTo(&actual))
	})

	t.Run("Too large for float64", func(t *testing.T) {
		tooLarge := decimal.RequireFromString("1e400")
		wrapped := NewDecimal(tooLarge)
		var actual float64
		require.Error(t, wrapped.UnwrapTo(&actual))
	})
}

func Test_DecimalUnwrapToFloat32(t *testing.T) {
	t.Run("Unwraps within range", func(t *testing.T) {
		expected := float32(1.33342)
		d := decimal.NewFromFloat(float64(expected))
		wrapped := NewDecimal(d)
		var actual float32
		err := wrapped.UnwrapTo(&actual)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Unwraps with alias type", func(t *testing.T) {
		type myfloat float32
		expected := myfloat(1.33342)
		d := decimal.NewFromFloat(float64(expected))
		wrapped := NewDecimal(d)
		var actual myfloat
		err := wrapped.UnwrapTo(&actual)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Unwraps with loss of precision", func(t *testing.T) {
		precisionLoss := decimal.RequireFromString("1.23456789012345678901234567890")
		wrapped := NewDecimal(precisionLoss)
		var actual float32
		err := wrapped.UnwrapTo(&actual)
		require.NoError(t, err)
		assert.Equal(t, float32(precisionLoss.InexactFloat64()), actual)
	})

	t.Run("Too small for float32", func(t *testing.T) {
		tooSmall := decimal.RequireFromString("1e-400") // smaller than ~5e-324
		wrapped := NewDecimal(tooSmall)
		var actual float32
		require.Error(t, wrapped.UnwrapTo(&actual))
	})

	t.Run("Too large for float32", func(t *testing.T) {
		tooLarge := decimal.RequireFromString("1e400")
		wrapped := NewDecimal(tooLarge)
		var actual float32
		require.Error(t, wrapped.UnwrapTo(&actual))
	})
}
