package arg_parser_config

import "unsafe"

// ArgumentsDescription contains specification of flag arguments
type ArgumentsDescription struct {
	amountType              ArgAmountType
	synopsisHelpDescription string
	defaultValues           []string
	allowedValues           map[string]bool
}

// ArgumentsDescriptionSrc contains source data for cast to ArgumentsDescription
type ArgumentsDescriptionSrc struct {
	AmountType              ArgAmountType
	SynopsisHelpDescription string
	DefaultValues           []string
	AllowedValues           map[string]bool
}

// CastPtr converts src to ArgumentsDescription pointer
func (src ArgumentsDescriptionSrc) CastPtr() *ArgumentsDescription {
	return (*ArgumentsDescription)(unsafe.Pointer(&src))
}

// GetAmountType - AmountType field getter
func (i *ArgumentsDescription) GetAmountType() ArgAmountType {
	if i == nil {
		return ArgAmountTypeNoArgs
	}
	return i.amountType
}

// GetSynopsisHelpDescription - SynopsisHelpDescription field getter
func (i *ArgumentsDescription) GetSynopsisHelpDescription() string {
	if i == nil {
		return ""
	}
	return i.synopsisHelpDescription
}

// GetDefaultValues - DefaultValues field getter
func (i *ArgumentsDescription) GetDefaultValues() []string {
	if i == nil {
		return nil
	}
	return i.defaultValues
}

// GetAllowedValues - AllowedValues field getter
func (i *ArgumentsDescription) GetAllowedValues() map[string]bool {
	if i == nil {
		return nil
	}
	return i.allowedValues
}
