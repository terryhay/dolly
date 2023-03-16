package page

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/tools/index"
	"github.com/terryhay/dolly/tools/size"
)

func TestBody(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *Body

		require.Equal(t, index.Zero, pointer.RowCount())
		require.Equal(t, Row{}, pointer.Row(index.RandIndex()))
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		t.Parallel()

		body := MakeBody(nil)
		require.Equal(t, index.Zero, body.RowCount())
		require.Equal(t, Row{}, body.Row(index.RandIndex()))

		row := MakeRow(size.WidthTab, MakeRowChunk(gofakeit.Color(), StyleTextDefault))
		body = MakeBody([]Row{row})

		require.Equal(t, index.MakeIndex(1), body.RowCount())
		require.Equal(t, row, body.Row(0))
	})
}
