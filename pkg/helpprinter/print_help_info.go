package helpprinter

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"strings"
)

// PrintHelpInfo prints help information by ArgParserConfig object
func PrintHelpInfo(config dollyconf.ArgParserConfig) {
	builder := strings.Builder{}

	builder.WriteString(CreateNameChapter(
		config.GetAppDescription().GetAppName(),
		config.GetAppDescription().GetNameHelpInfo()))

	builder.WriteString(CreateSynopsisChapter(
		config.GetAppDescription().GetAppName(),
		config.GetNamelessCommandDescription(),
		config.GetCommandDescriptions(),
		config.GetFlagDescriptions()))

	builder.WriteString(CreateDescriptionChapter(
		config.GetAppDescription().GetDescriptionHelpInfo(),
		config.GetNamelessCommandDescription(),
		config.GetCommandDescriptions(),
		config.GetFlagDescriptions()))

	fmt.Println(builder.String())
}
