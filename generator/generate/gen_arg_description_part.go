package generate

import (
	"fmt"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"strings"
)

const (
	argumentsDescriptionNilPart = `
%[1]s%[2]s nil`

	argumentsDescriptionPrefix = `
%[1]s%[2]sapConf.ArgumentsDescriptionSrc{
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
%[1]s	}.CastPtr(),`
	argumentsDescriptionPostfix = `
%[1]s}.CastPtr()`
)

// GenArgDescriptionPart - creates a paste part with argument description
func GenArgDescriptionPart(
	argumentsDescription *confYML.ArgumentsDescription,
	indent string,
	pasteArgDescriptionPrefix bool,
) string {

	prefix := ""
	if pasteArgDescriptionPrefix {
		prefix = "ArgDescription: "
	}

	if argumentsDescription == nil {
		return fmt.Sprintf(argumentsDescriptionNilPart, indent, prefix)
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

func getArgAmountTypeElement(argAmountType apConf.ArgAmountType) string {
	argAmountTypeElement := "apConf.ArgAmountTypeNoArgs"
	switch argAmountType {
	case apConf.ArgAmountTypeSingle:
		argAmountTypeElement = "apConf.ArgAmountTypeSingle"
	case apConf.ArgAmountTypeList:
		argAmountTypeElement = "apConf.ArgAmountTypeList"
	}
	return argAmountTypeElement
}
