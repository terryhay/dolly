package config_yaml

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyconf"
)

// ArgumentsDescription - description of command line flag arguments
type ArgumentsDescription struct {
	AmountType              dollyconf.ArgAmountType
	SynopsisHelpDescription string

	// optional
	DefaultValues []string
	AllowedValues []string
}

// GetAmountType - AmountType field getter
func (i *ArgumentsDescription) GetAmountType() dollyconf.ArgAmountType {
	if i == nil {
		return dollyconf.ArgAmountTypeNoArgs
	}
	return i.AmountType
}

// GetSynopsisHelpDescription - SynopsisHelpDescription field getter
func (i *ArgumentsDescription) GetSynopsisHelpDescription() string {
	if i == nil {
		return ""
	}
	return i.SynopsisHelpDescription
}

// GetDefaultValues - DefaultValues field getter
func (i *ArgumentsDescription) GetDefaultValues() []string {
	if i == nil {
		return nil
	}
	return i.DefaultValues
}

// GetAllowedValues - AllowedValues field getter
func (i *ArgumentsDescription) GetAllowedValues() []string {
	if i == nil {
		return nil
	}
	return i.AllowedValues
}

type argumentsDescriptionsSource struct {
	AmountType          string `yaml:"amount_type"`
	SynopsisDescription string `yaml:"synopsis_description"`

	DefaultValues []string `yaml:"default_values"`
	AllowedValues []string `yaml:"allowed_values"`
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (i *ArgumentsDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	source := new(argumentsDescriptionsSource)
	if err = unmarshal(&source); err != nil {
		return err
	}

	if len(source.AmountType) == 0 {
		return fmt.Errorf(`argumentsDescriptions unmarshal error: no required field "amount_type"`)
	}
	i.AmountType, err = argAmountStrType2argAmountType(source.AmountType)
	if err != nil {
		return fmt.Errorf(`argumentsDescriptions unmarshal error: can't convert string value "amount_type": %v`, err)
	}

	if len(source.SynopsisDescription) == 0 {
		return fmt.Errorf(`argumentsDescriptions unmarshal error: no required field "synopsis_description"`)
	}
	i.SynopsisHelpDescription = source.SynopsisDescription

	// don't check optional fields
	i.DefaultValues = source.DefaultValues
	i.AllowedValues = source.AllowedValues

	return nil
}

func argAmountStrType2argAmountType(argAmountStrType string) (dollyconf.ArgAmountType, error) {
	switch argAmountStrType {
	case "single":
		return dollyconf.ArgAmountTypeSingle, nil
	case "list":
		return dollyconf.ArgAmountTypeList, nil
	default:
		return dollyconf.ArgAmountTypeNoArgs,
			fmt.Errorf(`unexpected "amount_type" value: %s\nallowed values: "single", "array"`, argAmountStrType)
	}
}
