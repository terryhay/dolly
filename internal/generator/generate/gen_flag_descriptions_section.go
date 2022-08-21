package generate

import (
	"fmt"
	"github.com/terryhay/dolly/internal/generator/config_yaml"
	"github.com/terryhay/dolly/internal/generator/id_template_data_creator"
	"strings"
)

const (
	flagDescriptionsNilPart = `
		// flagDescriptions
		nil`

	flagDescriptionMapPrefix = `
		// flagDescriptions
		map[dollyconf.Flag]*dollyconf.FlagDescription{`
	flagDescriptionMapElementKeyPart = `
			%s: {`
	flagDescriptionMapElementDescriptionHelpInfo = `
				DescriptionHelpInfo:  "%s",`
	flagDescriptionMapElementPostfix = `
			},`
	flagDescriptionMapPostfix = `
		}`
)

// FlagDescriptionsSection - string with flag constant definitions list
type FlagDescriptionsSection string

// GenFlagDescriptionsSection creates a paste section with flag descriptions
func GenFlagDescriptionsSection(
	flagDescriptions []*config_yaml.FlagDescription,
	flagsIDTemplateData map[string]*id_template_data_creator.IDTemplateData) FlagDescriptionsSection {

	if len(flagDescriptions) == 0 {
		return flagDescriptionsNilPart
	}

	builder := new(strings.Builder)
	builder.WriteString(flagDescriptionMapPrefix)

	var (
		flagDescription      *config_yaml.FlagDescription
		argumentsDescription *config_yaml.ArgumentsDescription
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

	return FlagDescriptionsSection(builder.String())
}
