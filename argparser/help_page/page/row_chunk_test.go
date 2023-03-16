package page

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
)

func TestRowChunk(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *RowChunk

		require.Equal(t, StyleTextDefault, pointer.GetStyle())
		require.Empty(t, pointer.GetText())
		require.Equal(t, size.WidthZero, pointer.CountRunes())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		t.Parallel()

		text := gofakeit.Color()
		rowChunk := MakeRowChunk(text, StyleTextBold)

		require.Equal(t, StyleTextBold, rowChunk.GetStyle())
		require.Equal(t, text, rowChunk.GetText())
		require.Equal(t, size.MakeWidth(len([]rune(text))), rowChunk.CountRunes())
	})
}

func TestCreateStyledText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string
		rowChunk RowChunk
		exp      string
	}{
		{
			caseName: "empty",
		},
		{
			caseName: "default_implicit",
			rowChunk: MakeRowChunk(coty.RandNameCommand().String()),
			exp:      coty.RandNameCommand().String(),
		},
		{
			caseName: "default",
			rowChunk: MakeRowChunk(coty.RandNameCommand().String(), StyleTextDefault),
			exp:      coty.RandNameCommand().String(),
		},
		{
			caseName: "bold",
			rowChunk: MakeRowChunk(coty.RandNameCommand().String(), StyleTextBold),
			exp:      textStyleBoldOpen + coty.RandNameCommand().String() + textStyleClose,
		},
		{
			caseName: "underlined",
			rowChunk: MakeRowChunk(coty.RandNameCommand().String(), StyleTextUnderlined),
			exp:      textStyleUnderlinedOpen + coty.RandNameCommand().String() + textStyleClose,
		},
		{
			caseName: "underlined_bold",
			rowChunk: MakeRowChunk(coty.RandNameCommand().String(), StyleTextBold, StyleTextUnderlined),
			exp:      textStyleBoldOpen + textStyleUnderlinedOpen + coty.RandNameCommand().String() + textStyleClose + textStyleClose,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.exp, CreateStyledText(tc.rowChunk))
			require.Equal(t, tc.rowChunk.GetText(), RemoveStyleTextMarkers(CreateStyledText(tc.rowChunk)))
		})
	}
}

func TestRemoveStyleTextMarkers(t *testing.T) {
	t.Parallel()

	s := RemoveStyleTextMarkers(`[1m<empty>[0m empty command description`)
	require.Equal(t, "<empty> empty command description", s)
}
