package generate

import (
	"fmt"
	"strings"

	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

func appendNameAppConstant(builder *strings.Builder, nameApp coty.NameApp) *strings.Builder {
	if builder == nil {
		return builder
	}

	builder = sbt.NewRow(builder, size.WidthZero, "const (")
	defer func() {
		builder = sbt.NewRow(builder, size.WidthZero, ")")
		builder = sbt.BreakRow(builder)
	}()

	builder = sbt.NewRow(builder, size.WidthTab, "// NameApp - name of the application")
	builder = sbt.NewRow(builder, size.WidthTab, fmt.Sprintf(`NameApp coty.NameApp = "%s"`, nameApp))

	return builder
}
