package row

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/man_style_help/runes"
	"github.com/terryhay/dolly/tools/size"
)

func TestRow(t *testing.T) {
	t.Parallel()

	t.Run("getters", func(t *testing.T) {
		shiftIndex := size.MakeWidth(gofakeit.Uint32())
		cells := []termbox.Cell{
			{
				Ch: runes.RuneColon,
			},
		}
		row := MakeRow(shiftIndex, cells)

		require.Equal(t, shiftIndex, row.GetShiftIndex())
		require.Equal(t, cells, row.GetCells())
	})

	t.Run("String", func(t *testing.T) {
		row := MakeRow(
			1,
			[]termbox.Cell{
				{
					Ch: runes.RuneColon,
				},
			})

		require.Equal(t, " :", row.String())
	})
}
