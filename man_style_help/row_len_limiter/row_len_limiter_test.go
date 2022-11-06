package row_len_limiter

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	minRowLenMin     = 10
	minRowLenOptimum = 14
)

func TestCalcRowLenLimitMin(t *testing.T) {
	t.Parallel()

	t.Run("min_value", func(t *testing.T) {
		require.Equal(t, minRowLenMin, calcRowLenLimitMin(TerminalMinWidth))
	})

	t.Run("max_value", func(t *testing.T) {
		require.Equal(t, defaultRowLenMin, calcRowLenLimitMin(defaultRowLenMax))
	})
}

func TestCalcRowLenLimitOptimum(t *testing.T) {
	t.Parallel()

	t.Run("min_value", func(t *testing.T) {
		require.Equal(t, minRowLenOptimum, calcRowLenLimitOptimum(TerminalMinWidth))
	})

	t.Run("max_value", func(t *testing.T) {
		require.Equal(t, defaultRowLenOptimum, calcRowLenLimitOptimum(defaultRowLenMax))
	})
}

func TestRowLenLimiter(t *testing.T) {
	t.Parallel()

	rll := MakeRowLenLimiter()

	defaultLimit := MakeDefaultRowLenLimit()
	assert.Equal(t, RowLenLimiter{
		lastTerminalWidth: defaultLimit.Max(),

		rowLenLimit:        defaultLimit,
		defaultRowLenLimit: defaultLimit,
	}, rll)

	assert.Equal(t, RowLenLimit{}, rll.GetRowLenLimit(0))
	assert.Equal(t, RowLenLimit{}, rll.GetRowLenLimit(0))
	assert.Equal(t, defaultLimit, rll.GetRowLenLimit(defaultLimit.Max()))
}
