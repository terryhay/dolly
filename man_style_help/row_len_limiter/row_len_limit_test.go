package row_len_limiter

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/tools/size"
)

func TestRowLenLimitGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer_getters", func(t *testing.T) {
		t.Parallel()
		var limit *RowLenLimit

		require.ErrorIs(t, limit.IsValid(), ErrIsValidMin)

		require.Equal(t, 0, limit.Min().Int())
		require.Equal(t, 0, limit.Optimum().Int())
		require.Equal(t, 0, limit.Max().Int())

		require.Equal(t, 0, len(limit.String()))
	})

	t.Run("initialized", func(t *testing.T) {
		t.Parallel()

		widthMin := 20
		widthOptimum := 30
		widthMax := 40

		limit := MakeRowLenLimit(size.MakeWidth(widthMin), size.MakeWidth(widthOptimum), size.MakeWidth(widthMax))
		require.NoError(t, limit.IsValid())

		require.Equal(t, widthMin, limit.Min().Int())
		require.Equal(t, widthOptimum, limit.Optimum().Int())
		require.Equal(t, widthMax, limit.Max().Int())

		require.Equal(t, fmt.Sprintf("[min: %d; optimum: %d; max: %d]", widthMin, widthOptimum, widthMax), limit.String())
	})
}

func TestRowLenLimitShifting(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer_shifting", func(t *testing.T) {
		var limit *RowLenLimit
		assert.Equal(t, RowLenLimit{}, limit.ApplyTabShift(size.MakeWidth(gofakeit.Uint8())))
	})

	t.Run("shifting", func(t *testing.T) {
		widthMin := 2 + int(gofakeit.Uint8())
		optimum := widthMin + int(gofakeit.Uint8())
		widthMax := optimum + int(gofakeit.Uint8())

		limit := MakeRowLenLimit(size.MakeWidth(widthMin), size.MakeWidth(optimum), size.MakeWidth(widthMax))
		shifted := limit.ApplyTabShift(1)

		require.Equal(t, limit.Min().Int()-1, shifted.Min().Int())
		require.Equal(t, limit.Optimum().Int()-1, shifted.Optimum().Int())
		require.Equal(t, limit.Max().Int()-1, shifted.Max().Int())
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

func TestClone(t *testing.T) {
	t.Parallel()

	{
		limit := MakeRowLenLimit(
			size.MakeWidth(gofakeit.Uint8()),
			size.MakeWidth(gofakeit.Uint8()),
			size.MakeWidth(gofakeit.Uint8()),
		)
		clone := CloneRowLenLimit(limit)

		require.NotNil(t, clone)
		require.Equal(t, limit.Min(), clone.Min())
		require.Equal(t, limit.Optimum(), clone.Optimum())
		require.Equal(t, limit.Max(), clone.Max())
	}
}

func TestIsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string
		rll      RowLenLimit
		expErr   error
	}{
		{
			caseName: "zero_min",
			rll:      MakeRowLenLimit(size.WidthZero, size.WidthTab, size.WidthTab),
			expErr:   ErrIsValidMin,
		},
		{
			caseName: "zero_optimum",
			rll:      MakeRowLenLimit(size.WidthTab, size.WidthZero, size.WidthTab),
			expErr:   ErrIsValidOptimum,
		},
		{
			caseName: "zero_max",
			rll:      MakeRowLenLimit(size.WidthTab, size.WidthTab, size.WidthZero),
			expErr:   ErrIsValidMax,
		},
		{
			caseName: "optimum_less_than_min",
			rll:      MakeRowLenLimit(size.MakeWidth(2), size.MakeWidth(1), size.MakeWidth(10)),
			expErr:   ErrIsValidMinMoreThanOptimum,
		},
		{
			caseName: "max_less_than_optimum",
			rll:      MakeRowLenLimit(size.MakeWidth(2), size.MakeWidth(11), size.MakeWidth(10)),
			expErr:   ErrIsValidOptimumMoreThanMax,
		},
		{
			caseName: "valid",
			rll:      MakeRowLenLimit(size.MakeWidth(2), size.MakeWidth(5), size.MakeWidth(10)),
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			if tc.expErr == nil {
				require.NoError(t, tc.rll.IsValid())
			}
			require.ErrorIs(t, tc.rll.IsValid(), tc.expErr)
		})
	}
}
