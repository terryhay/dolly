package page

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/tools/size"
)

func TestRow(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *Row

		require.Empty(t, pointer.GetTextStyled())
		require.Equal(t, size.WidthZero, pointer.GetTextRuneCount())
		require.Equal(t, size.WidthZero, pointer.GetMarginLeft())
	})

	t.Run("initialized", func(t *testing.T) {
		t.Parallel()

		row := Row{}
		require.Empty(t, row.GetTextStyled())
		require.Equal(t, size.WidthZero, row.GetTextRuneCount())
		require.Equal(t, size.WidthZero, row.GetMarginLeft())

		row = MakeRow(size.WidthZero)
		require.Empty(t, row.GetTextStyled())
		require.Equal(t, size.WidthZero, row.GetTextRuneCount())
		require.Equal(t, size.WidthZero, row.GetMarginLeft())

		row = MakeRow(size.WidthTab,
			MakeRowChunk("<empty>", StyleTextBold),
			MakeRowChunkSpaces(1),
			MakeRowChunk("empty command description"),
		)
		require.Equal(t, `[1m<empty>[0m empty command description`, row.GetTextStyled())
		require.Equal(t, size.MakeWidth(len("<empty> empty command description")), row.GetTextRuneCount())
		require.Equal(t, size.WidthTab, row.GetMarginLeft())
	})
}
