package generate

import (
	"fmt"
	"github.com/terryhay/dolly/internal/generator/id_template_data_creator"
	"strings"
)

const (
	flagConstantsPrefixPart = `
const (`
	flagConstantsFirstLinePart = `
	// %s - %s
	%s dollyconf.Flag = "%s"`
	flagConstantsLinePart = `
	// %s - %s
	%s = "%s"`
	flagConstantsPostfixPart = `
)`
)

// FlagStringIDListSection - string with flag constant definitions list
type FlagStringIDListSection string

// GenFlagIDConstantsSection - creates a paste section flag constants
func GenFlagIDConstantsSection(
	flagsTemplateData map[string]*id_template_data_creator.IDTemplateData,
) FlagStringIDListSection {

	if len(flagsTemplateData) == 0 {
		return ""
	}

	builder := strings.Builder{}
	builder.WriteString(flagConstantsPrefixPart)

	sortedFlagsTemplateData := sortByNameID(flagsTemplateData)

	templateData := sortedFlagsTemplateData[0]
	builder.WriteString(fmt.Sprintf(flagConstantsFirstLinePart,
		templateData.GetNameID(),
		templateData.GetComment(),
		templateData.GetNameID(),
		templateData.GetCallName()))

	for i := 1; i < len(sortedFlagsTemplateData); i++ {
		templateData = sortedFlagsTemplateData[i]
		builder.WriteString(fmt.Sprintf(flagConstantsLinePart,
			templateData.GetNameID(),
			templateData.GetComment(),
			templateData.GetNameID(),
			templateData.GetCallName()))
	}

	builder.WriteString(flagConstantsPostfixPart)
	return FlagStringIDListSection(builder.String())
}
