package string_to_cells

import (
	"testing"

	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/require"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
)

func TestForeground(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *foreground

		require.Equal(t, termbox.Attribute(0), pointer.Get())
		require.NotPanics(t, func() { pointer.Set(termbox.AttrBold) })
		require.NotPanics(t, func() { pointer.DropBack() })
	})

	t.Run("initialized", func(t *testing.T) {
		t.Parallel()

		fg := makeForeground(2)

		fg.Set(termbox.AttrBold)
		require.Equal(t, termbox.AttrBold, fg.Get())

		fg.Set(termbox.AttrUnderline)
		require.Equal(t, termbox.AttrBold|termbox.AttrUnderline, fg.Get())

		fg.DropBack()
		require.Equal(t, termbox.AttrBold, fg.Get())

		fg.DropBack()
		require.Equal(t, termbox.Attribute(0), fg.Get())
	})
}

func TestStringToCells(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string
		text     string
		exp      []termbox.Cell
	}{
		{
			caseName: "empty",
		},
		{
			caseName: "unstyled",
			text:     "cmd",
			exp: []termbox.Cell{
				{Ch: 'c'},
				{Ch: 'm'},
				{Ch: 'd'},
			},
		},
		{
			caseName: "enexpected_style",
			text:     "\u001B[5mcmd\u001B[0m",
			exp: []termbox.Cell{
				{Ch: '['},
				{Ch: '5'},
				{Ch: 'm'},
				{Ch: 'c'},
				{Ch: 'm'},
				{Ch: 'd'},
			},
		},
		{
			caseName: "bold",
			text:     hp.CreateStyledText(hp.MakeRowChunk("cmd", hp.StyleTextBold)),
			exp: []termbox.Cell{
				{Ch: 'c', Fg: termbox.AttrBold},
				{Ch: 'm', Fg: termbox.AttrBold},
				{Ch: 'd', Fg: termbox.AttrBold},
			},
		},
		{
			caseName: "underlined",
			text:     hp.CreateStyledText(hp.MakeRowChunk("cmd", hp.StyleTextUnderlined)),
			exp: []termbox.Cell{
				{Ch: 'c', Fg: termbox.AttrUnderline},
				{Ch: 'm', Fg: termbox.AttrUnderline},
				{Ch: 'd', Fg: termbox.AttrUnderline},
			},
		},
		{
			caseName: "bold_underlined",
			text:     hp.CreateStyledText(hp.MakeRowChunk("cmd", hp.StyleTextBold, hp.StyleTextUnderlined)),
			exp: []termbox.Cell{
				{Ch: 'c', Fg: termbox.AttrBold | termbox.AttrUnderline},
				{Ch: 'm', Fg: termbox.AttrBold | termbox.AttrUnderline},
				{Ch: 'd', Fg: termbox.AttrBold | termbox.AttrUnderline},
			},
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.exp, StringToCells(tc.text))
		})
	}
}
