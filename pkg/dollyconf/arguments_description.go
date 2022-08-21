package dollyconf

// ArgumentsDescription contains specification of flag arguments
type ArgumentsDescription struct {
	AmountType              ArgAmountType
	SynopsisHelpDescription string
	DefaultValues           []string
	AllowedValues           map[string]bool
}

// GetAmountType - AmountType field getter
func (i *ArgumentsDescription) GetAmountType() ArgAmountType {
	if i == nil {
		return ArgAmountTypeNoArgs
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
func (i *ArgumentsDescription) GetAllowedValues() map[string]bool {
	if i == nil {
		return nil
	}
	return i.AllowedValues
}
