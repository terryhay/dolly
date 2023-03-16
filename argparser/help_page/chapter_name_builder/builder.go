package chapter_name_builder

import (
	"fmt"

	hp "github.com/terryhay/dolly/argparser/help_page/page"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
)

// AppendRows creates chapter NAME AppendRows
func AppendRows(rows []hp.Row, appName coty.NameApp, info coty.InfoChapterNAME) []hp.Row {

	if len(appName) == 0 || len(info) == 0 {
		return rows
	}

	return append(rows,
		hp.MakeRow(size.WidthZero,
			hp.MakeRowChunk(`NAME`, hp.StyleTextBold)),
		hp.MakeRow(size.WidthTab,
			hp.MakeRowChunk(appName.String(), hp.StyleTextBold),
			hp.MakeRowChunk(fmt.Sprintf(` – %s`, info)),
		),
	)
}
