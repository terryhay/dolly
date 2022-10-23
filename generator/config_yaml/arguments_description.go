package config_yaml

import (
	"fmt"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
)

// ArgumentsDescription - description of command line flag arguments
type ArgumentsDescription struct {
	AmountType              apConf.ArgAmountType
	SynopsisHelpDescription string

	// optional
	DefaultValues []string
	AllowedValues []string
}

// GetAmountType - AmountType field getter
func (ad *ArgumentsDescription) GetAmountType() apConf.ArgAmountType {
	if ad == nil {
		return apConf.ArgAmountTypeNoArgs
	}
	return ad.AmountType
}

// GetSynopsisHelpDescription - SynopsisHelpDescription field getter
func (ad *ArgumentsDescription) GetSynopsisHelpDescription() string {
	if ad == nil {
		return ""
	}
	return ad.SynopsisHelpDescription
}

// GetDefaultValues - DefaultValues field getter
func (ad *ArgumentsDescription) GetDefaultValues() []string {
	if ad == nil {
		return nil
	}
	return ad.DefaultValues
}

// GetAllowedValues - AllowedValues field getter
func (ad *ArgumentsDescription) GetAllowedValues() []string {
	if ad == nil {
		return nil
	}
	return ad.AllowedValues
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (ad *ArgumentsDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = ad

	proxy := struct {
		AmountType          string `yaml:"amount_type"`
		SynopsisDescription string `yaml:"synopsis_description"`

		DefaultValues []string `yaml:"default_values"`
		AllowedValues []string `yaml:"allowed_values"`
	}{}
	if err = unmarshal(&proxy); err != nil {
		return err
	}

	if len(proxy.AmountType) == 0 {
		return fmt.Errorf(`argumentsDescriptions unmarshal error: no required field "amount_type"`)
	}
	ad.AmountType, err = argAmountStrType2argAmountType(proxy.AmountType)
	if err != nil {
		return fmt.Errorf(`argumentsDescriptions unmarshal error: can't convert string value "amount_type": %s`, err.Error())
	}

	if len(proxy.SynopsisDescription) == 0 {
		return fmt.Errorf(`argumentsDescriptions unmarshal error: no required field "synopsis_description"`)
	}
	ad.SynopsisHelpDescription = proxy.SynopsisDescription

	// don't check optional fields
	ad.DefaultValues = proxy.DefaultValues
	ad.AllowedValues = proxy.AllowedValues

	return nil
}

func argAmountStrType2argAmountType(argAmountStrType string) (apConf.ArgAmountType, error) {
	switch argAmountStrType {
	case "single":
		return apConf.ArgAmountTypeSingle, nil
	case "list":
		return apConf.ArgAmountTypeList, nil
	default:
		return apConf.ArgAmountTypeNoArgs,
			fmt.Errorf(`unexpected "amount_type" value: %s\nallowed values: "single", "array"`, argAmountStrType)
	}
}
