package chapter_name_builder

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	coty "github.com/terryhay/dolly/tools/common_types"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

func TestAppendRows(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string

		appName coty.NameApp
		info    coty.InfoChapterNAME

		expStyled   string
		expUnstyled string
	}{
		{
			caseName: "no_data",
		},
		{
			caseName: "only_app_name",
			appName:  coty.RandNameApp(),
		},
		{
			caseName: "only_info",
			info:     coty.RandInfoChapterName(),
		},
		{
			caseName: "common",
			appName:  coty.RandNameApp(),
			info:     coty.RandInfoChapterName(),
			expStyled: fmt.Sprintf(`[1mNAME[0m
    [1m%s[0m – %s`,
				coty.RandNameApp(),
				coty.RandInfoChapterName(),
			),
			expUnstyled: fmt.Sprintf(`NAME
    %s – %s`,
				coty.RandNameApp(),
				coty.RandInfoChapterName(),
			),
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			rows := AppendRows(make([]hp.Row, 0), tc.appName, tc.info)

			builder := &strings.Builder{}
			for i := range rows {
				if i > 0 {
					builder = sbt.BreakRow(builder)
				}

				tab := hp.MakeRowChunkSpaces(rows[i].GetMarginLeft())
				builder = sbt.Append(builder, tab.GetText(), rows[i].GetTextStyled())
			}

			str := builder.String()
			require.Equal(t, tc.expStyled, str)
			require.Equal(t, tc.expUnstyled, hp.RemoveStyleTextMarkers(str))
		})
	}
}
