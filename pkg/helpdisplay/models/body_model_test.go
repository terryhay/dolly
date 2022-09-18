package models

import (
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/helpdisplay/data"
	rll "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"testing"
)

func TestBodyModelErrors(t *testing.T) {
	t.Parallel()

	bm := NewBodyModel(data.Page{}, TerminalSize{})

	{
		prm, err := bm.GetParagraph(0)
		require.Nil(t, prm)
		require.Nil(t, err)
	}

	{
		err := bm.Update(TerminalSize{}, 0)
		require.NotNil(t, err)
	}
}

func TestBodyModelShifting(t *testing.T) {
	t.Parallel()

	terminalSize := MakeTerminalSize(
		rll.MakeDefaultRowLenLimit(),
		10,
	)
	bm := NewBodyModel(data.Page{
		Paragraphs: []*data.Paragraph{
			{
				Text: "1",
			},
			{
				Text: "2",
			},
		},
	},
		terminalSize,
	)

	require.Nil(t, bm.Shift(terminalSize.GetHeight(), 2))
	require.Nil(t, bm.Shift(terminalSize.GetHeight(), -2))
}
