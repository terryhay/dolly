package generate

import (
	"strings"

	ce "github.com/terryhay/dolly/generator/config_entity"
	"github.com/terryhay/dolly/tools/size"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

// appendPlaceholderConstants creates a paste section with command ids
func appendPlaceholderConstants(builder *strings.Builder, genDataPlaceholdersSorted []*ce.GenComponents) *strings.Builder {
	if builder == nil || len(genDataPlaceholdersSorted) == 0 {
		return builder
	}

	builder = sbt.NewRow(builder, size.WidthZero, "const (")
	defer func() {
		builder = sbt.NewRow(builder, size.WidthZero, ")")
		builder = sbt.BreakRow(builder)
	}()

	genData := genDataPlaceholdersSorted[0]
	builder = sbt.NewRow(builder, size.WidthTab, "// ", genData.GetNameID(), " - ", genData.GetComment())
	builder = sbt.NewRow(builder, size.WidthTab, genData.GetNameID(), ` coty.IDPlaceholder = iota + 1`)

	for i := 1; i < len(genDataPlaceholdersSorted); i++ {
		genData = genDataPlaceholdersSorted[i]

		builder = sbt.BreakRow(builder)
		builder = sbt.NewRow(builder, size.WidthTab, "// ", genData.GetNameID(), " - ", genData.GetComment())
		builder = sbt.NewRow(builder, size.WidthTab, genData.GetNameID())
	}

	return builder
}
