package header_model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/man_style_help/row"
	rowLenLimitMock "github.com/terryhay/dolly/man_style_help/row_len_limiter/mock"
	s2cells "github.com/terryhay/dolly/man_style_help/string_to_cells"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestHeaderModelGetters(t *testing.T) {
	t.Parallel()

	var hm *HeaderModel
	hm.Update(ts.TerminalSize{})
	require.Equal(t, row.Row{}, hm.GetViewRow())

	appName := coty.RandNameApp()
	hm = NewHeaderModel(appName, ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitForTerminalWidth25(), 20))
	require.Equal(t, appName.String(), strings.TrimSpace(hm.GetViewRow().String()))
	hm.Update(ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitForTerminalWidth25(), 20))
}

func TestMakeHeaderModel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string

		headerText coty.NameApp
		size       ts.TerminalSize

		expHeaderModel *HeaderModel
	}{
		{
			caseName:       "empty_headerText_and_rowLenLimit",
			expHeaderModel: &HeaderModel{},
		},
		{
			caseName: "empty_headerText_with_max_rowLenLimit",
			size:     ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMax(), 20),

			expHeaderModel: &HeaderModel{
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitMax(),
				shift:            33,
			},
		},
		{
			caseName:   "header_with_max_rowLenLimit",
			headerText: "example",
			size:       ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMax(), 20),

			expHeaderModel: &HeaderModel{
				cellsSrc:         s2cells.StringToCells("example"),
				cellsOut:         s2cells.StringToCells("example"),
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitMax(),
				shift:            29,
			},
		},
		{
			caseName:   "header_with_rowLenLimit25",
			headerText: "example",
			size:       ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitForTerminalWidth25(), 20),

			expHeaderModel: &HeaderModel{
				cellsSrc:         s2cells.StringToCells("example"),
				cellsOut:         s2cells.StringToCells("example"),
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),
				shift:            5,
			},
		},
		{
			caseName:   "header_with_min_rowLenLimit",
			headerText: "example",
			size:       ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMin(), 20),

			expHeaderModel: &HeaderModel{
				cellsSrc:         s2cells.StringToCells("example"),
				cellsOut:         s2cells.StringToCells("example"),
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitMin(),
				shift:            3,
			},
		},
		{
			caseName:   "using_max_terminal_size",
			headerText: "Help info: example utility",
			size:       ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitForTerminalWidth25(), 20),

			expHeaderModel: &HeaderModel{
				cellsSrc:         s2cells.StringToCells("Help info: example utility"),
				cellsOut:         s2cells.StringToCells("Help info: example utility"),
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),
			},
		},
		{
			caseName:   "header_cutting",
			headerText: "Help info: example application",
			size:       ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMin(), 20),

			expHeaderModel: &HeaderModel{
				cellsSrc:         s2cells.StringToCells("Help info: example application"),
				cellsOut:         s2cells.StringToCells("Help info: exampl..."),
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitMin(),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseName, func(t *testing.T) {
			headerModel := NewHeaderModel(tc.headerText, tc.size)
			require.NotNil(t, headerModel)
			require.Equal(t, tc.expHeaderModel, headerModel)
		})
	}
}
