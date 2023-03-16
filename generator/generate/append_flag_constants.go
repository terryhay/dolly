package generate

import (
	"strings"

	ce "github.com/terryhay/dolly/generator/config_entity"
	"github.com/terryhay/dolly/tools/size"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

// appendFlagConstants creates a paste section flag constants
func appendFlagConstants(builder *strings.Builder, genCompFlagsSorted []*ce.GenComponents) *strings.Builder {
	if builder == nil || len(genCompFlagsSorted) == 0 {
		return builder
	}

	builder = sbt.NewRow(builder, size.WidthZero, "const (")
	defer func() {
		builder = sbt.NewRow(builder, size.WidthZero, ")")
		builder = sbt.BreakRow(builder)
	}()

	genComp := genCompFlagsSorted[0]
	builder = sbt.NewRow(builder, size.WidthTab, "// ", genComp.GetNameID(), " - ", genComp.GetComment())
	builder = sbt.NewRow(builder, size.WidthTab, genComp.GetNameID(), ` coty.NameFlag = "`, genComp.GetName().String(), `"`)

	for i := 1; i < len(genCompFlagsSorted); i++ {
		genComp = genCompFlagsSorted[i]

		builder = sbt.BreakRow(builder)
		builder = sbt.NewRow(builder, size.WidthTab, "// ", genComp.GetNameID(), " - ", genComp.GetComment())
		builder = sbt.NewRow(builder, size.WidthTab, genComp.GetNameID(), ` coty.NameFlag = "`, genComp.GetName().String(), `"`)
	}

	return builder
}
