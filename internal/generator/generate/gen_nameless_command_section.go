package generate

import (
	"fmt"
	"github.com/terryhay/dolly/internal/generator/config_yaml"
	"github.com/terryhay/dolly/internal/generator/id_template_data_creator"
	"strings"
)

const (
	namelessCommandDescriptionNilPart = `
		// namelessCommandDescription
		nil`

	namelessCommandDescriptionPrefixPart = `
		// namelessCommandDescription
		dollyconf.NewNamelessCommandDescription(`
	namelessCommandDescriptionCommandIDPart = `
			%s,`
	namelessCommandDescriptionDescriptionHelpInfoPart = `
			"%s",`
	namelessCommandDescriptionFlagMapNilPart = `
			nil,`
	namelessCommandDescriptionFlagMapPrefixPart = `
			map[dollyconf.Flag]bool{`
	namelessCommandDescriptionFlagMapLinePart = `
				%s: true,`
	namelessCommandDescriptionFlagMapPostfixPart = `
			},`
	namelessCommandDescriptionPostfixPart = `
		)`
)

// NamelessCommandDescriptionSection - string with nameless command description section
type NamelessCommandDescriptionSection string

// GenNamelessCommandComponent creates a paste section with nameless command description
func GenNamelessCommandComponent(
	namelessCommandDescription *config_yaml.NamelessCommandDescription,
	namelessCommandIDTemplateData *id_template_data_creator.IDTemplateData,
	flagsIDTemplateData map[string]*id_template_data_creator.IDTemplateData,
) NamelessCommandDescriptionSection {

	if namelessCommandDescription == nil {
		return namelessCommandDescriptionNilPart
	}

	builder := strings.Builder{}
	builder.WriteString(
		namelessCommandDescriptionPrefixPart)
	builder.WriteString(
		fmt.Sprintf(namelessCommandDescriptionCommandIDPart, namelessCommandIDTemplateData.GetID()))
	builder.WriteString(
		fmt.Sprintf(namelessCommandDescriptionDescriptionHelpInfoPart, namelessCommandDescription.GetDescriptionHelpInfo()))
	builder.WriteString(
		GenArgDescriptionPart(namelessCommandDescription.GetArgumentsDescription(), "\t\t\t", false) + ",")
	builder.WriteString(
		createNamelessCommandDescriptionFlagsPart(namelessCommandDescription.GetRequiredFlags(), flagsIDTemplateData))
	builder.WriteString(
		createNamelessCommandDescriptionFlagsPart(namelessCommandDescription.GetOptionalFlags(), flagsIDTemplateData))
	builder.WriteString(
		namelessCommandDescriptionPostfixPart)

	return NamelessCommandDescriptionSection(builder.String())
}

func createNamelessCommandDescriptionFlagsPart(
	flags []string,
	flagsIDTemplateData map[string]*id_template_data_creator.IDTemplateData,
) string {

	if len(flags) == 0 {
		return namelessCommandDescriptionFlagMapNilPart
	}

	builder := strings.Builder{}
	builder.WriteString(namelessCommandDescriptionFlagMapPrefixPart)

	for _, flag := range flags {
		builder.WriteString(
			fmt.Sprintf(namelessCommandDescriptionFlagMapLinePart, flagsIDTemplateData[flag].GetNameID()))
	}
	builder.WriteString(namelessCommandDescriptionFlagMapPostfixPart)

	return builder.String()
}
