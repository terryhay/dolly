package plain_help_out

import (
	"fmt"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"strings"
)

// PrintHelpInfo prints help information by argParserConfig object
func PrintHelpInfo(config apConf.ArgParserConfig) {
	builder := strings.Builder{}

	builder.WriteString(createNameChapter(
		config.GetAppDescription().GetAppName(),
		config.GetAppDescription().GetNameHelpInfo()))

	builder.WriteString(createSynopsisChapter(
		config.GetAppDescription().GetAppName(),
		config.GetNamelessCommandDescription(),
		config.GetCommandDescriptions(),
		config.GetFlagDescriptions()))

	builder.WriteString(createDescriptionChapter(
		config.GetAppDescription().GetDescriptionHelpInfo(),
		config.GetNamelessCommandDescription(),
		config.GetCommandDescriptions(),
		config.GetFlagDescriptions()))

	fmt.Println(builder.String())
}
