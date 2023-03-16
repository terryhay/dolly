package chapter_synopsis_builder

import (
	"fmt"
	"sort"
	"strings"

	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
)

// AppendRows creates chapter SYNOPSIS
func AppendRows(rows []hp.Row, appName coty.NameApp, commands []*apConf.Command) []hp.Row {
	if len(appName) == 0 || len(commands) == 0 {
		return rows
	}

	rows = append(rows, hp.Row{}, hp.MakeRow(size.WidthZero, hp.MakeRowChunk(`SYNOPSIS`, hp.StyleTextBold)))
	rows = appendCommandRows(rows, appName, commands)

	return rows
}

func appendCommandRows(rows []hp.Row, appName coty.NameApp, commands []*apConf.Command) []hp.Row {
	rowCommands := make([]hp.Row, 0, len(commands))

	const countChunksExpected = 8
	for _, command := range commands {
		chunks := make([]hp.RowChunk, 0, countChunksExpected)

		switch {
		case len(command.GetNameMain()) > 0:
			chunks = append(chunks,
				hp.MakeRowChunk(
					fmt.Sprintf("%s %s ", appName, command.CreateStringWithCommandNames()),
					hp.StyleTextBold,
				),
			)

		default:
			chunks = append(chunks, hp.MakeRowChunk(fmt.Sprintf("%s ", appName), hp.StyleTextBold))
		}

		if len(command.GetPlaceholders()) == 0 {
			rowCommands = append(rowCommands, hp.MakeRow(size.WidthTab, chunks...))
			continue
		}

		for i := 0; i < len(command.GetPlaceholders()); i++ {
			if i > 0 {
				chunks = append(chunks, hp.MakeRowChunkSpaces(1))
			}
			chunks = appendChunksPlaceholder(chunks, command.GetPlaceholders()[i])
		}

		rowCommands = append(rowCommands, hp.MakeRow(size.WidthTab, chunks...))
	}

	sort.Slice(rowCommands, func(l, r int) bool {
		return rowCommands[l].GetTextStyled() < rowCommands[r].GetTextStyled()
	})

	return append(rows, rowCommands...)
}

// appendChunksPlaceholder creates command options info string from command placeholders
func appendChunksPlaceholder(chunks []hp.RowChunk, placeholder *apConf.Placeholder) []hp.RowChunk {
	flags := placeholder.GetDescriptionFlags()

	if placeholder.GetIsFlagOptional() {
		chunks = append(chunks, hp.MakeRowChunk(`[`))
	}

	chunks = appendChunksFlags(chunks, flags)
	chunks = appendChunksArgument(chunks, placeholder.GetArgument(), func() bool {
		var nameFlag coty.NameFlag
		for nameFlag = range flags {
			break
		}

		return nameFlag.IsDoubleDash()
	}())

	if placeholder.GetIsFlagOptional() {
		chunks = append(chunks, hp.MakeRowChunk(`]`))
	}

	return chunks
}

func appendChunksFlags(chunks []hp.RowChunk, flagsByNames map[coty.NameFlag]*apConf.Flag) []hp.RowChunk {
	flagNames := make([]string, 0, len(flagsByNames))
	oneLetterFlags := true
	for _, flag := range flagsByNames {
		if len(flag.GetNameMain()) > 2 {
			oneLetterFlags = false
		}

		flagNames = append(flagNames, flag.GetNameMain().String())
	}

	sort.Slice(flagNames, func(l, r int) bool {
		lName, rName := strings.ToLower(flagNames[l]), strings.ToLower(flagNames[r])
		if lName == rName {
			return flagNames[l] < flagNames[r]
		}

		return lName < rName
	})

	const countOneLetterFlagsWithoutCompressionMax = 3
	if oneLetterFlags && len(flagNames) > countOneLetterFlagsWithoutCompressionMax {
		builder := &strings.Builder{}
		builder.WriteRune('-')
		for _, name := range flagNames {
			builder.WriteRune(rune(name[1]))
		}
		return append(chunks, hp.MakeRowChunk(builder.String(), hp.StyleTextBold))
	}

	for i, name := range flagNames {
		if i > 0 {
			const separator = " | "
			chunks = append(chunks, hp.MakeRowChunk(separator))
		}
		chunks = append(chunks, hp.MakeRowChunk(name, hp.StyleTextBold))
	}

	return chunks
}

// chunksArguments creates argument options info chunks from placeholder arguments description
func appendChunksArgument(chunks []hp.RowChunk, argument *apConf.Argument, isDoubleDashFlag bool) []hp.RowChunk {
	if argument == nil {
		return chunks
	}

	switch {
	case argument.GetIsOptional() && isDoubleDashFlag:
		chunks = append(chunks, hp.MakeRowChunk("[="))

	case argument.GetIsOptional():
		chunks = append(chunks, hp.MakeRowChunk("["))

	case isDoubleDashFlag:
		chunks = append(chunks, hp.MakeRowChunk("="))

	default:
		chunks = append(chunks, hp.MakeRowChunk(" "))
	}

	chunks = append(chunks, hp.MakeRowChunk(argument.GetSynopsisHelpDescription(), hp.StyleTextUnderlined))

	switch {
	case argument.GetIsOptional() && argument.GetIsList():
		chunks = append(chunks,
			hp.MakeRowChunk(" "),
			hp.MakeRowChunk("...", hp.StyleTextUnderlined),
			hp.MakeRowChunk("]"),
		)

	case argument.GetIsOptional():
		chunks = append(chunks, hp.MakeRowChunk("]"))

	case argument.GetIsList():
		chunks = append(chunks,
			hp.MakeRowChunk(" "),
			hp.MakeRowChunk("...", hp.StyleTextUnderlined),
		)
	}

	return chunks
}
