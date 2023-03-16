package chapter_description_builder

import (
	"sort"

	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
)

// AppendRows creates chapter DESCRIPTION
func AppendRows(
	rows []hp.Row,
	textsIntroduction []coty.InfoChapterDESCRIPTION,
	commands []*apConf.Command,
) []hp.Row {

	if len(textsIntroduction) == 0 &&
		len(commands) == 0 {

		return rows
	}

	const countRowsExpected = 16

	rowsDESCRIPTION := make([]hp.Row, 0, countRowsExpected)
	rowsDESCRIPTION = appendIntroductionRows(rowsDESCRIPTION, textsIntroduction)
	rowsDESCRIPTION = appendCommandsRows(rowsDESCRIPTION, commands)
	rowsDESCRIPTION = appendFlagRows(rowsDESCRIPTION, commands)

	if len(rowsDESCRIPTION) == 0 {
		return rows
	}

	rows = append(rows, hp.Row{}, makeRowDESCRIPTION())
	rows = append(rows, rowsDESCRIPTION...)

	return rows
}

func makeRowDESCRIPTION() hp.Row {
	const text = "DESCRIPTION"
	return hp.MakeRow(size.WidthZero, hp.MakeRowChunk(text, hp.StyleTextBold))
}

func appendIntroductionRows(
	rows []hp.Row,
	introduction []coty.InfoChapterDESCRIPTION,
) []hp.Row {

	for _, text := range introduction {
		rows = append(rows,
			hp.MakeRow(size.WidthTab, hp.MakeRowChunk(text.String())),
			hp.MakeRow(size.WidthZero),
		)
	}

	return rows
}

func appendCommandsRows(rows []hp.Row, commands []*apConf.Command) []hp.Row {
	if len(commands) == 1 && len(commands[0].GetNameMain()) == 0 {
		// Don't show help info if there is only nameless command
		return rows
	}

	rows = append(rows, MakeRowTheCommandsAreAsFollows())

	shift := 0
	if len(commands) > 0 && len(commands[0].GetNameMain()) == 0 {
		shift++

		chunkTitle := hp.MakeRowChunk("<empty>", hp.StyleTextBold)
		rows = append(rows,
			hp.MakeRow(size.WidthTab,
				chunkTitle,
				hp.MakeRowChunkSpaces(size.Dif(size.WidthDescriptionColumnShift, chunkTitle.CountRunes())),
				hp.MakeRowChunk(commands[0].GetDescriptionHelpInfo()),
			),
			hp.Row{},
		)
	}

	for i := shift; i < len(commands); i++ {
		chunkTitle := hp.MakeRowChunk(commands[i].CreateStringWithCommandNames(), hp.StyleTextBold)
		widthsDif := size.Dif(size.WidthDescriptionColumnShift, chunkTitle.CountRunes())

		if widthsDif == size.WidthZero {
			rows = append(rows,
				hp.MakeRow(size.WidthTab, chunkTitle),
				hp.MakeRow(size.WidthTab+size.WidthDescriptionColumnShift, hp.MakeRowChunk(commands[i].GetDescriptionHelpInfo())),
				hp.Row{},
			)

			continue
		}

		rows = append(rows,
			hp.MakeRow(size.WidthTab,
				chunkTitle,
				hp.MakeRowChunkSpaces(widthsDif),
				hp.MakeRowChunk(commands[i].GetDescriptionHelpInfo()),
			),
			hp.Row{},
		)
	}

	return rows
}

// MakeRowTheCommandsAreAsFollows builds Row with text 'The commands are as follows:'
func MakeRowTheCommandsAreAsFollows() hp.Row {
	const text = "The commands are as follows:"
	return hp.MakeRow(size.WidthZero, hp.MakeRowChunk(text))
}

func appendFlagRows(rows []hp.Row, commands []*apConf.Command) []hp.Row {
	type FlagWithArg struct {
		Flag *apConf.Flag
		Arg  *apConf.Argument
	}

	flagsWithArgsSorted := make([]FlagWithArg, 0, len(commands)*5)
	for _, command := range commands {
		for _, placeholder := range command.GetPlaceholders() {
			for _, flag := range placeholder.GetDescriptionFlags() {
				flagsWithArgsSorted = append(flagsWithArgsSorted, FlagWithArg{
					Flag: flag,
					Arg:  placeholder.GetArgument(),
				})
			}
		}
	}

	sort.Slice(flagsWithArgsSorted, func(l, r int) bool {
		return flagsWithArgsSorted[l].Flag.GetNameMain() < flagsWithArgsSorted[r].Flag.GetNameMain()
	})

	if len(flagsWithArgsSorted) == 0 {
		return rows
	}
	rows = append(rows, MakeRowTheFlagsAreAsFollows())

	for _, flagWithArg := range flagsWithArgsSorted {
		chunksTitle, runeCount := titleFlagRowChunks(flagWithArg.Flag, flagWithArg.Arg)
		widthsDif := size.Dif(size.WidthDescriptionColumnShift, runeCount)

		if widthsDif == size.WidthZero {
			rows = append(rows,
				hp.MakeRow(size.WidthTab, chunksTitle...),
				hp.MakeRow(size.WidthTab+size.WidthDescriptionColumnShift, hp.MakeRowChunk(flagWithArg.Flag.GetDescriptionHelpInfo())),
				hp.Row{},
			)
			continue
		}

		chunksTitle = append(chunksTitle,
			hp.MakeRowChunkSpaces(widthsDif),
			hp.MakeRowChunk(flagWithArg.Flag.GetDescriptionHelpInfo()),
		)
		rows = append(rows, hp.MakeRow(size.WidthTab, chunksTitle...), hp.Row{})
	}

	return rows
}

// MakeRowTheFlagsAreAsFollows builds Row with text 'The flags are as follows:'
func MakeRowTheFlagsAreAsFollows() hp.Row {
	const text = "The flags are as follows:"
	return hp.MakeRow(size.WidthZero, hp.MakeRowChunk(text))
}

func titleFlagRowChunks(flag *apConf.Flag, argument *apConf.Argument) ([]hp.RowChunk, size.Width) {
	chunks := make([]hp.RowChunk, 0)

	chunks = append(chunks, hp.MakeRowChunk(flag.GetNameMain().String(), hp.StyleTextBold))
	chunks = appendChunksArguments(chunks, argument, flag.GetNameMain().IsDoubleDash())

	if len(flag.GetNamesAdditional()) > 0 {
		additionalFlagsSorted := make([]coty.NameFlag, 0, len(flag.GetNamesAdditional()))
		for name := range flag.GetNamesAdditional() {
			additionalFlagsSorted = append(additionalFlagsSorted, name)
		}
		sort.Slice(additionalFlagsSorted, func(l, r int) bool {
			return additionalFlagsSorted[l] < additionalFlagsSorted[r]
		})

		for name := range flag.GetNamesAdditional() {
			chunks = append(chunks,
				hp.MakeRowChunk(", "),
				hp.MakeRowChunk(name.String(), hp.StyleTextBold),
			)

			chunks = appendChunksArguments(chunks, argument, name.IsDoubleDash())
		}
	}

	runeCount := size.WidthZero
	for _, chunk := range chunks {
		runeCount += chunk.CountRunes()
	}

	return chunks, runeCount
}

// chunksArguments creates argument options info chunks from placeholder arguments description
func appendChunksArguments(chunks []hp.RowChunk, argument *apConf.Argument, isDoubleDashFlag bool) []hp.RowChunk {
	if argument == nil {
		return chunks
	}

	var prefix string
	switch {
	case argument.GetIsOptional() && isDoubleDashFlag:
		prefix = "[="

	case argument.GetIsOptional():
		prefix = "["

	case isDoubleDashFlag:
		prefix = "="

	default:
		prefix = " "
	}

	var postfix string
	switch {
	case argument.GetIsOptional() && argument.GetIsList():
		postfix = "] ..."

	case argument.GetIsOptional():
		postfix = "]"

	case argument.GetIsList():
		postfix = " ..."
	}

	return append(chunks,
		hp.MakeRowChunk(prefix),
		hp.MakeRowChunk(argument.GetSynopsisHelpDescription(), hp.StyleTextUnderlined),
		hp.MakeRowChunk(postfix),
	)
}
