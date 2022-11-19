package config_yaml

import (
	"fmt"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"unsafe"
)

// ArgumentsDescription - description of command line flag arguments
type ArgumentsDescription struct {
	amountType              apConf.ArgAmountType
	synopsisHelpDescription string

	// optional
	defaultValues []string
	allowedValues []string
}

// GetAmountType - amountType field getter
func (ad *ArgumentsDescription) GetAmountType() apConf.ArgAmountType {
	if ad == nil {
		return apConf.ArgAmountTypeNoArgs
	}
	return ad.amountType
}

// GetSynopsisHelpDescription - synopsisHelpDescription field getter
func (ad *ArgumentsDescription) GetSynopsisHelpDescription() string {
	if ad == nil {
		return ""
	}
	return ad.synopsisHelpDescription
}

// GetDefaultValues - defaultValues field getter
func (ad *ArgumentsDescription) GetDefaultValues() []string {
	if ad == nil {
		return nil
	}
	return ad.defaultValues
}

// GetAllowedValues - allowedValues field getter
func (ad *ArgumentsDescription) GetAllowedValues() []string {
	if ad == nil {
		return nil
	}
	return ad.allowedValues
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
	ad.amountType, err = argAmountStrType2argAmountType(proxy.AmountType)
	if err != nil {
		return fmt.Errorf(`argumentsDescriptions unmarshal error: can't convert string value "amount_type": %s`, err.Error())
	}

	if len(proxy.SynopsisDescription) == 0 {
		return fmt.Errorf(`argumentsDescriptions unmarshal error: no required field "synopsis_description"`)
	}
	ad.synopsisHelpDescription = proxy.SynopsisDescription

	// don't check optional fields
	ad.defaultValues = proxy.DefaultValues
	ad.allowedValues = proxy.AllowedValues

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

// ArgumentsDescriptionSrc - source for construct a description of command line flag arguments
type ArgumentsDescriptionSrc struct {
	AmountType              apConf.ArgAmountType
	SynopsisHelpDescription string

	// optional
	DefaultValues []string
	AllowedValues []string
}

// ToConstPtr converts src to AppHelpDescription pointer
func (m ArgumentsDescriptionSrc) ToConstPtr() *ArgumentsDescription {
	return (*ArgumentsDescription)(unsafe.Pointer(&m))
}
