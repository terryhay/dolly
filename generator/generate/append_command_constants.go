package generate

import (
	"strings"

	ce "github.com/terryhay/dolly/generator/config_entity"
	"github.com/terryhay/dolly/tools/size"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

// appendCommandConstants creates a paste section with command ids
func appendCommandConstants(builder *strings.Builder, genDataCommandsSorted []*ce.GenComponents) *strings.Builder {
	if builder == nil || len(genDataCommandsSorted) == 0 {
		return builder
	}

	builder = sbt.NewRow(builder, size.WidthZero, "const (")
	defer func() {
		builder = sbt.NewRow(builder, size.WidthZero, ")")
		builder = sbt.BreakRow(builder)
	}()

	genData := genDataCommandsSorted[0]
	builder = sbt.NewRow(builder, size.WidthTab, "// ", genData.GetNameID(), " - ", genData.GetComment())
	builder = sbt.NewRow(builder, size.WidthTab, genData.GetNameID(), ` coty.NameCommand = "`, genData.GetName().String(), `"`)

	for i := 1; i < len(genDataCommandsSorted); i++ {
		genData = genDataCommandsSorted[i]

		builder = sbt.BreakRow(builder)
		builder = sbt.NewRow(builder, size.WidthTab, "// ", genData.GetNameID(), " - ", genData.GetComment())
		builder = sbt.NewRow(builder, size.WidthTab, genData.GetNameID(), ` coty.NameCommand = "`, genData.GetName().String(), `"`)
	}

	return builder
}
