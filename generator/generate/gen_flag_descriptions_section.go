package generate

import (
	"fmt"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	creator "github.com/terryhay/dolly/generator/id_template_data_creator"
	"strings"
)

const (
	flagDescriptionsNilPart = `
		FlagDescriptionSlice: nil`

	flagDescriptionSlicePrefix = `
		FlagDescriptionSlice: []*apConf.FlagDescription{`
	flagDescriptionSliceFlagsPart = `
				Flags: []apConf.Flag{
					%s,
				},`
	flagDescriptionSliceElementKeyPart = `
			apConf.FlagDescriptionSrc{`
	flagDescriptionSliceElementDescriptionHelpInfo = `
				DescriptionHelpInfo:  "%s",`
	flagDescriptionSliceElementPostfix = `
			}.ToConstPtr(),`
	flagDescriptionSlicePostfix = `
		}`
)

// sectionFlagDescriptions - string with flag constant definitions list
type sectionFlagDescriptions string

// genFlagDescriptionsSection creates a paste section with flag descriptions
func genFlagDescriptionsSection(
	flagDescriptions []*confYML.FlagDescription,
	flagsIDTemplateData map[string]*creator.IDTemplateData) sectionFlagDescriptions {

	if len(flagDescriptions) == 0 {
		return flagDescriptionsNilPart
	}

	builder := new(strings.Builder)
	builder.WriteString(flagDescriptionSlicePrefix)

	for _, descFlag := range flagDescriptions {
		builder.WriteString(flagDescriptionSliceElementKeyPart)

		flags := descFlag.ExtractSortedFlags()
		flagIDs := make([]string, 0, len(flags))
		for _, flag := range flags {
			if templateData, contain := flagsIDTemplateData[flag]; contain {
				flagIDs = append(flagIDs, templateData.GetNameID())
			}
		}
		builder.WriteString(fmt.Sprintf(flagDescriptionSliceFlagsPart, strings.Join(flagIDs, ",\t\t\t")))

		builder.WriteString(fmt.Sprintf(flagDescriptionSliceElementDescriptionHelpInfo,
			descFlag.GetDescriptionHelpInfo()))

		if descArg := descFlag.GetArgumentsDescription(); descArg != nil {
			builder.WriteString(GenArgDescriptionPart(descArg, "\t\t\t\t", true))
			builder.WriteString(",")
		}

		builder.WriteString(flagDescriptionSliceElementPostfix)
	}
	builder.WriteString(flagDescriptionSlicePostfix)

	return sectionFlagDescriptions(builder.String())
}
