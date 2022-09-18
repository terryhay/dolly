package models

import (
	"github.com/stretchr/testify/require"
	rowLenLimitMock "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter/mock"
	"testing"
)

func TestMakeHeaderModel(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string

		headerText string
		size       TerminalSize

		expectedHeaderModel *HeaderModel
	}{
		{
			caseName:            "empty_headerText_and_rowLenLimit",
			expectedHeaderModel: NewHeaderModel("", TerminalSize{}),
		},
		{
			caseName: "empty_headerText_with_max_rowLenLimit",
			size:     MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMax(), 20),

			expectedHeaderModel: &HeaderModel{
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitMax(),
				shift:           33,
			},
		},
		{
			caseName:   "header_with_max_rowLenLimit",
			headerText: "example",
			size:       MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMax(), 20),

			expectedHeaderModel: &HeaderModel{
				headerCells:     textToCells("example"),
				outputCells:     textToCells("example"),
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitMax(),
				shift:           29,
			},
		},
		{
			caseName:   "header_with_rowLenLimit25",
			headerText: "example",
			size:       MakeTerminalSize(rowLenLimitMock.GetRowLenLimitForTerminalWidth25(), 20),

			expectedHeaderModel: &HeaderModel{
				headerCells:     textToCells("example"),
				outputCells:     textToCells("example"),
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),
				shift:           5,
			},
		},
		{
			caseName:   "header_with_min_rowLenLimit",
			headerText: "example",
			size:       MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMin(), 20),

			expectedHeaderModel: &HeaderModel{
				headerCells:     textToCells("example"),
				outputCells:     textToCells("example"),
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitMin(),
				shift:           3,
			},
		},
		{
			caseName:   "using_max_terminal_size",
			headerText: "Help info: example utility",
			size:       MakeTerminalSize(rowLenLimitMock.GetRowLenLimitForTerminalWidth25(), 20),

			expectedHeaderModel: &HeaderModel{
				headerCells:     textToCells("Help info: example utility"),
				outputCells:     textToCells("Help info: example utility"),
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),
			},
		},
		{
			caseName:   "header_cutting",
			headerText: "Help info: example application",
			size:       MakeTerminalSize(rowLenLimitMock.GetRowLenLimitMin(), 20),

			expectedHeaderModel: &HeaderModel{
				headerCells:     textToCells("Help info: example application"),
				outputCells:     textToCells("Help info: exampl..."),
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitMin(),
			},
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			headerModel := NewHeaderModel(td.headerText, td.size)
			require.Equal(t, td.expectedHeaderModel, headerModel)
		})
	}
}
