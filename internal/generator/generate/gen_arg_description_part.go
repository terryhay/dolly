package generate

import (
	"fmt"
	"github.com/terryhay/dolly/internal/generator/config_yaml"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"strings"
)

const (
	argumentsDescriptionNilPart = `
%[1]s%[2]s nil`

	argumentsDescriptionPrefix = `
%[1]s%[2]s&dollyconf.ArgumentsDescription{
%[1]s	AmountType:              %[3]s,
%[1]s	SynopsisHelpDescription: "%[4]s",`
	argumentsDescriptionDefaultValuesPrefix = `
%[1]s	DefaultValues: []string{`
	argumentsDescriptionAllowedValuesPrefix = `
%[1]s	AllowedValues: map[string]bool{`
	argumentsDescriptionVariantValue = `
%[1]s		"%[2]s",`
	argumentsDescriptionMapVariantValue = `
%[1]s		"%[2]s": true,`
	argumentsDescriptionVariantValuesPostfix = `
%[1]s	},`
	argumentsDescriptionPostfix = `
%[1]s}`
)

// GenArgDescriptionPart - creates a paste part with argument description
func GenArgDescriptionPart(
	argumentsDescription *config_yaml.ArgumentsDescription,
	indent string,
	pasteArgDescriptionPrefix bool,
) string {

	prefix := ""
	if pasteArgDescriptionPrefix {
		prefix = "ArgDescription: "
	}

	if argumentsDescription == nil {
		return fmt.Sprintf(fmt.Sprintf(argumentsDescriptionNilPart, indent, prefix))
	}

	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf(argumentsDescriptionPrefix,
		indent,
		prefix,
		getArgAmountTypeElement(argumentsDescription.GetAmountType()),
		argumentsDescription.GetSynopsisHelpDescription()))

	if defaultValues := argumentsDescription.GetDefaultValues(); len(defaultValues) > 0 {
		builder.WriteString(fmt.Sprintf(argumentsDescriptionDefaultValuesPrefix, indent))
		for _, value := range defaultValues {
			builder.WriteString(fmt.Sprintf(argumentsDescriptionVariantValue, indent, value))
		}
		builder.WriteString(fmt.Sprintf(argumentsDescriptionVariantValuesPostfix, indent))
	}

	if allowedValues := argumentsDescription.GetAllowedValues(); len(allowedValues) > 0 {
		builder.WriteString(fmt.Sprintf(argumentsDescriptionAllowedValuesPrefix, indent))
		for _, value := range allowedValues {
			builder.WriteString(fmt.Sprintf(argumentsDescriptionMapVariantValue, indent, value))
		}
		builder.WriteString(fmt.Sprintf(argumentsDescriptionVariantValuesPostfix, indent))
	}

	builder.WriteString(fmt.Sprintf(argumentsDescriptionPostfix, indent))

	return builder.String()
}

func getArgAmountTypeElement(argAmountType dollyconf.ArgAmountType) string {
	argAmountTypeElement := "dollyconf.ArgAmountTypeNoArgs"
	switch argAmountType {
	case dollyconf.ArgAmountTypeSingle:
		argAmountTypeElement = "dollyconf.ArgAmountTypeSingle"
	case dollyconf.ArgAmountTypeList:
		argAmountTypeElement = "dollyconf.ArgAmountTypeList"
	}
	return argAmountTypeElement
}
