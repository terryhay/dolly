package generate

import (
	"strings"

	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

func appendChapterDescriptionInfo(builder *strings.Builder, marginLeft size.Width, infoChapterDescription []coty.InfoChapterDESCRIPTION) *strings.Builder {
	if builder == nil || len(infoChapterDescription) == 0 {
		return builder
	}

	builder = sbt.NewRow(builder, marginLeft, "HelpInfoChapterDESCRIPTION: []string{")
	defer func() { builder = sbt.NewRow(builder, marginLeft, "}") }()

	for _, info := range infoChapterDescription {
		builder = sbt.NewRow(builder, marginLeft+size.WidthTab, `"`, info.String(), `"`)
	}

	return builder
}
