package header_model

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/argparser/helpdisplay/page"
	"github.com/terryhay/dolly/argparser/helpdisplay/row"
	rowLenLimitMock "github.com/terryhay/dolly/argparser/helpdisplay/row_len_limiter/mock"
	ts "github.com/terryhay/dolly/argparser/helpdisplay/terminal_size"
	"strings"
	"testing"
)

func TestHeaderModelGetters(t *testing.T) {
	t.Parallel()

	var hm *HeaderModel
	err := hm.Update(ts.TerminalSize{})
	require.Nil(t, err)
	require.Equal(t, row.Row{}, hm.GetViewRow())

	header := gofakeit.Name()
	hm, err = NewHeaderModel(page.MakeParagraph(0, header), ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitForTerminalWidth25(), 20))
	require.Nil(t, err)
	require.Equal(t, header, strings.TrimSpace(hm.GetViewRow().String()))
	require.Nil(t, hm.Update(ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitForTerminalWidth25(), 20)))
}

func TestMakeHeaderModel(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string

		headerText string
		size       ts.TerminalSize

		expectedHeaderModel *HeaderModel
		expectedErr         bool
	}{
		{
			caseName:    "empty_headerText_and_rowLenLimit",
			expectedErr: true,
		},
		{
			caseName: "empty_headerText_with_max_rowLenLimit",
			size:     ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMax(), 20),

			expectedHeaderModel: &HeaderModel{
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitMax(),
				shift:            33,
			},
		},
		{
			caseName:   "header_with_max_rowLenLimit",
			headerText: "example",
			size:       ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMax(), 20),

			expectedHeaderModel: &HeaderModel{
				paragraph:        page.MakeParagraph(0, "example"),
				outputCells:      page.TextToCells("example"),
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitMax(),
				shift:            29,
			},
		},
		{
			caseName:   "header_with_rowLenLimit25",
			headerText: "example",
			size:       ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitForTerminalWidth25(), 20),

			expectedHeaderModel: &HeaderModel{
				paragraph:        page.MakeParagraph(0, "example"),
				outputCells:      page.TextToCells("example"),
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),
				shift:            5,
			},
		},
		{
			caseName:   "header_with_min_rowLenLimit",
			headerText: "example",
			size:       ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMin(), 20),

			expectedHeaderModel: &HeaderModel{
				paragraph:        page.MakeParagraph(0, "example"),
				outputCells:      page.TextToCells("example"),
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitMin(),
				shift:            3,
			},
		},
		{
			caseName:   "using_max_terminal_size",
			headerText: "Help info: example utility",
			size:       ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitForTerminalWidth25(), 20),

			expectedHeaderModel: &HeaderModel{
				paragraph:        page.MakeParagraph(0, "Help info: example utility"),
				outputCells:      page.TextToCells("Help info: example utility"),
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),
			},
		},
		{
			caseName:   "header_cutting",
			headerText: "Help info: example application",
			size:       ts.MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMin(), 20),

			expectedHeaderModel: &HeaderModel{
				paragraph:        page.MakeParagraph(0, "Help info: example application"),
				outputCells:      page.TextToCells("Help info: exampl..."),
				usingRowLenLimit: rowLenLimitMock.GetRowLenLimitMin(),
			},
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			headerModel, err := NewHeaderModel(page.MakeParagraph(0, td.headerText), td.size)
			if td.expectedErr {
				require.Nil(t, headerModel)
				require.NotNil(t, err)
				return
			}

			require.Equal(t, td.expectedHeaderModel, headerModel)
			require.Nil(t, err)
		})
	}
}
