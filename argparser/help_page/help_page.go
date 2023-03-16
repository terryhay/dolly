package help_page

import (
	"sort"

	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	builderDESCRIPTION "github.com/terryhay/dolly/argparser/help_page/chapter_description_builder"
	builderNAME "github.com/terryhay/dolly/argparser/help_page/chapter_name_builder"
	builderSYNOPSIS "github.com/terryhay/dolly/argparser/help_page/chapter_synopsis_builder"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
)

// MakeBody creates Body object for display help info
func MakeBody(config apConf.ArgParserConfig) hp.Body {
	commands := prepareCommands(config.GetCommandNameless(), config.GetCommands())

	const countRowsExpected = 89
	rows := make([]hp.Row, 0, countRowsExpected)

	rows = builderNAME.AppendRows(rows,
		config.GetAppDescription().GetAppName(),
		config.GetAppDescription().GetNameHelpInfo(),
	)

	rows = builderSYNOPSIS.AppendRows(rows,
		config.GetAppDescription().GetAppName(),
		commands,
	)

	rows = builderDESCRIPTION.AppendRows(rows,
		config.GetHelpInfoChapterDESCRIPTION(),
		commands,
	)

	return hp.MakeBody(rows)
}

func prepareCommands(commandNameless *apConf.Command, commands []*apConf.Command) []*apConf.Command {
	res := make([]*apConf.Command, 0, 1+len(commands))

	if commandNameless != nil {
		res = append(res, commandNameless)
	}

	res = append(res, commands...)

	sort.Slice(res, func(l, r int) bool {
		return res[l].GetNameMain() < res[r].GetNameMain()
	})

	return res
}
