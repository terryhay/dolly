package generate

import (
	"strings"

	ce "github.com/terryhay/dolly/generator/config_entity"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

func appendPlaceholders(
	builder *strings.Builder,
	marginLeft size.Width,
	configEntity ce.ConfigEntity,
	namePlaceholders []coty.NamePlaceholder,
) *strings.Builder {
	if builder == nil || len(namePlaceholders) == 0 {
		return builder
	}

	builder = sbt.NewRow(builder, marginLeft, "Placeholders: []*apConf.PlaceholderOpt{")
	defer func() { builder = sbt.NewRow(builder, marginLeft, "},") }()

	marginArray := marginLeft + size.WidthTab
	for _, namePlaceholder := range namePlaceholders {
		placeholder := configEntity.PlaceholderByName(namePlaceholder)
		genDataPlaceholder := configEntity.GenCompPlaceholderByName(namePlaceholder)

		builder = sbt.NewRow(builder, marginArray, "{")
		{
			marginBody := marginArray + size.WidthTab
			builder = sbt.NewRow(builder, marginBody, "ID: ", genDataPlaceholder.GetNameID(), ",")
			builder = appendFlags(builder, marginBody, configEntity, placeholder.GetFlags())
			builder = appendArg(builder, marginBody, placeholder.GetArgument())
		}
		builder = sbt.NewRow(builder, marginArray, "},")
	}

	return builder
}

func appendFlags(
	builder *strings.Builder,
	marginLeft size.Width,
	configEntity ce.ConfigEntity,
	flags []*confYML.Flag,
) *strings.Builder {
	if builder == nil || len(flags) == 0 {
		return builder
	}

	builder = sbt.NewRow(builder, marginLeft, "FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{")
	defer func() { builder = sbt.NewRow(builder, marginLeft, "},") }()

	marginArray := marginLeft + size.WidthTab
	for _, flag := range flags {
		genDataFlag := configEntity.GenCompFlagByName(flag.GetMainName())
		builder = sbt.NewRow(builder, marginArray, genDataFlag.GetNameID(), `: {`)
		{
			marginBody := marginArray + size.WidthTab
			builder = sbt.NewRow(builder, marginBody, "NameMain: ", genDataFlag.GetNameID(), ",")
			if len(flag.GetAdditionalNames()) > 0 {
				builder = sbt.NewRow(builder, marginBody, `NamesAdditional:      map[coty.NameFlag]struct{}{`)

				marginBodySecond := marginBody + size.WidthTab
				for _, name := range flag.GetAdditionalNames() {
					genDataFlag = configEntity.GenCompFlagByName(name)
					builder = sbt.NewRow(builder, marginBodySecond, genDataFlag.GetNameID(), ",")
				}

				builder = sbt.NewRow(builder, marginBody, "},")
			}

			builder = sbt.NewRow(builder, marginBody, `HelpInfo: "`, flag.GetDescriptionHelpInfo().String(), `",`)
			if flag.GetIsOptional() {
				builder = sbt.NewRow(builder, marginBody, `IsOptional:          true`)
			}
		}
		builder = sbt.NewRow(builder, marginArray, "},")
	}

	return builder
}

func appendArg(builder *strings.Builder, marginLeft size.Width, arg *confYML.Argument) *strings.Builder {
	if builder == nil || arg == nil {
		return builder
	}

	builder = sbt.NewRow(builder, marginLeft, "Argument: &apConf.ArgumentOpt{")
	defer func() { builder = sbt.NewRow(builder, marginLeft, "},") }()

	marginBody := marginLeft + size.WidthTab
	if arg.GetIsList() {
		builder = sbt.NewRow(builder, marginBody, "IsList: coty.ArgAmountTypeList,")
	}

	marginBodySecond := marginBody + size.WidthTab

	if len(arg.GetDefaultValues()) > 0 {
		builder = sbt.NewRow(builder, marginBody, "DefaultValues: []string{")
		{
			for _, v := range arg.GetDefaultValues() {
				builder = sbt.NewRow(builder, marginBodySecond, `"`, v.String(), `",`)
			}
		}
		builder = sbt.NewRow(builder, marginBody, "},")
	}

	if len(arg.GetAllowedValues()) > 0 {
		builder = sbt.NewRow(builder, marginBody, "AllowedValues: []string{")
		{
			for _, v := range arg.GetAllowedValues() {
				builder = sbt.NewRow(builder, marginBodySecond, `"`, v.String(), `",`)
			}
		}
		builder = sbt.NewRow(builder, marginBody, "},")
	}

	if arg.GetIsOptional() {
		builder = sbt.NewRow(builder, marginBody, `IsOptional:          true,`)
	}

	return builder
}
