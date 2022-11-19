package generate

import (
	"fmt"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	creator "github.com/terryhay/dolly/generator/id_template_data_creator"
	"strings"
)

const (
	flagDescriptionsNilPart = `
		FlagDescriptions: nil`

	flagDescriptionMapPrefix = `
		FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{`
	flagDescriptionMapElementKeyPart = `
			%s: apConf.FlagDescriptionSrc{`
	flagDescriptionMapElementDescriptionHelpInfo = `
				DescriptionHelpInfo:  "%s",`
	flagDescriptionMapElementPostfix = `
			}.ToConstPtr(),`
	flagDescriptionMapPostfix = `
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
	builder.WriteString(flagDescriptionMapPrefix)

	var (
		flagDescription      *confYML.FlagDescription
		argumentsDescription *confYML.ArgumentsDescription
	)

	for i := 0; i < len(flagDescriptions); i++ {
		flagDescription = flagDescriptions[i]

		builder.WriteString(fmt.Sprintf(flagDescriptionMapElementKeyPart,
			flagsIDTemplateData[flagDescription.GetFlag()].GetNameID()))
		builder.WriteString(fmt.Sprintf(flagDescriptionMapElementDescriptionHelpInfo,
			flagDescription.GetDescriptionHelpInfo()))

		if argumentsDescription = flagDescription.GetArgumentsDescription(); argumentsDescription != nil {
			builder.WriteString(GenArgDescriptionPart(argumentsDescription, "\t\t\t\t", true))
			builder.WriteString(",")
		}

		builder.WriteString(flagDescriptionMapElementPostfix)
	}
	builder.WriteString(flagDescriptionMapPostfix)

	return sectionFlagDescriptions(builder.String())
}
