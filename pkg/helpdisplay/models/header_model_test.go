package models

import (
	"github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	rowLenLimitMock "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter/mock"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeHeaderModel(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string

		headerText  string
		rowLenLimit row_len_limiter.RowLenLimit

		expectedHeaderModel HeaderModel
	}{
		{
			caseName: "empty_headerText_and_rowLenLimit",
		},
		{
			caseName:    "empty_headerText_with_max_rowLenLimit",
			rowLenLimit: rowLenLimitMock.GetRowLenLimitMax(),

			expectedHeaderModel: HeaderModel{
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitMax(),
				shift:           33,
			},
		},
		{
			caseName:    "header_with_max_rowLenLimit",
			headerText:  "example",
			rowLenLimit: rowLenLimitMock.GetRowLenLimitMax(),

			expectedHeaderModel: HeaderModel{
				headerCells:     textToCells("example"),
				outputCells:     textToCells("example"),
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitMax(),
				shift:           29,
			},
		},
		{
			caseName:    "header_with_rowLenLimit25",
			headerText:  "example",
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),

			expectedHeaderModel: HeaderModel{
				headerCells:     textToCells("example"),
				outputCells:     textToCells("example"),
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),
				shift:           5,
			},
		},
		{
			caseName:    "header_with_min_rowLenLimit",
			headerText:  "example",
			rowLenLimit: rowLenLimitMock.GetRowLenLimitMin(),

			expectedHeaderModel: HeaderModel{
				headerCells:     textToCells("example"),
				outputCells:     textToCells("example"),
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitMin(),
				shift:           3,
			},
		},
		{
			caseName:    "using_max_terminal_size",
			headerText:  "Help info: example utility",
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),

			expectedHeaderModel: HeaderModel{
				headerCells:     textToCells("Help info: example utility"),
				outputCells:     textToCells("Help info: example utility"),
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),
			},
		},
		{
			caseName:    "header_cutting",
			headerText:  "Help info: example application",
			rowLenLimit: rowLenLimitMock.GetRowLenLimitMin(),

			expectedHeaderModel: HeaderModel{
				headerCells:     textToCells("Help info: example application"),
				outputCells:     textToCells("Help info: exampl..."),
				usingRowLeLimit: rowLenLimitMock.GetRowLenLimitMin(),
			},
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			headerModel := MakeHeaderModel(td.headerText, td.rowLenLimit)
			require.Equal(t, td.expectedHeaderModel, headerModel)
		})
	}
}
