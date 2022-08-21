package row_len_limiter

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRowLenLimitGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer_getters", func(t *testing.T) {
		var limit *RowLenLimit

		require.False(t, limit.IsValid())

		require.Equal(t, 0, limit.Min().ToInt())
		require.Equal(t, 0, limit.Optimum().ToInt())
		require.Equal(t, 0, limit.Max().ToInt())
	})

	t.Run("getters", func(t *testing.T) {
		min := 20
		optimum := 30
		max := 40

		limit := MakeRowLenLimit(RowSize(min), RowSize(optimum), RowSize(max))
		require.True(t, limit.IsValid())

		require.Equal(t, min, limit.Min().ToInt())
		require.Equal(t, optimum, limit.Optimum().ToInt())
		require.Equal(t, max, limit.Max().ToInt())
	})
}

func TestRowLenLimitShifting(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer_shifting", func(t *testing.T) {
		var limit *RowLenLimit
		assert.Equal(t, RowLenLimit{}, limit.ApplyTabShift(RowSize(gofakeit.Uint8())))
	})

	t.Run("shifting", func(t *testing.T) {
		min := 2 + int(gofakeit.Uint8())
		optimum := min + int(gofakeit.Uint8())
		max := optimum + int(gofakeit.Uint8())

		limit := MakeRowLenLimit(RowSize(min), RowSize(optimum), RowSize(max))
		shifted := limit.ApplyTabShift(1)

		require.Equal(t, limit.Min().ToInt()-1, shifted.Min().ToInt())
		require.Equal(t, limit.Optimum().ToInt()-1, shifted.Optimum().ToInt())
		require.Equal(t, limit.Max().ToInt()-1, shifted.Max().ToInt())
	})

	t.Run("extremal_shifting", func(t *testing.T) {
		limit := MakeRowLenLimit(
			defaultRowLenMin,
			defaultRowLenMax,
			defaultRowLenOptimum,
		)
		shifted := limit.ApplyTabShift(defaultRowLenMax)

		require.Equal(t, limit, shifted)
	})
}
