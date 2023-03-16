package config_yaml

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// Argument - description of mainName line flag arguments
type Argument struct {
	helpName coty.NameArgHelp

	// optional
	isList        bool
	defaultValues []coty.ArgValue
	allowedValues []coty.ArgValue
	isOptional    bool
}

// GetHelpName gets helpName field
func (a *Argument) GetHelpName() coty.NameArgHelp {
	if a == nil {
		return coty.NameArgHelpUndefined
	}
	return a.helpName
}

// GetIsList gets isList field
func (a *Argument) GetIsList() bool {
	if a == nil {
		return false
	}
	return a.isList
}

// GetDefaultValues gets defaultValues field
func (a *Argument) GetDefaultValues() []coty.ArgValue {
	if a == nil {
		return nil
	}
	return a.defaultValues
}

// GetAllowedValues gets allowedValues field
func (a *Argument) GetAllowedValues() []coty.ArgValue {
	if a == nil {
		return nil
	}
	return a.allowedValues
}

// GetIsOptional gets isOptional field
func (a *Argument) GetIsOptional() bool {
	if a == nil {
		return false
	}
	return a.isOptional
}

var (
	// ErrArgumentHelpName - no 'help_name' filed in source yaml file
	ErrArgumentHelpName = errors.New(`Argument.IsValid: 'arg_parser.placeholders[].argument.help_name'`)

	// ErrArgumentStopCharacter - unexpected argument value in default or allowed lists
	ErrArgumentStopCharacter = errors.New(`invalid object 'argument': unexpected character in some argument value`)

	// ErrArgumentDefaultValueIsNotAllowed - 'default_values' field value must be in 'allowed_values' field list
	ErrArgumentDefaultValueIsNotAllowed = errors.New(`invalid object 'argument': 'default_values' field value must be in 'allowed_values' field list`)
)

var stopCharsPattern = regexp.MustCompile(`[^\w\d\\/\.-_:"']`)

// IsValid check if Argument is valid
func (a *Argument) IsValid() error {
	if a == nil {
		return nil
	}

	if err := a.helpName.IsValid(); err != nil {
		return errors.Join(ErrArgumentHelpName, err)
	}

	for _, valAllowed := range a.allowedValues {
		if stopChar := stopCharsPattern.FindString(valAllowed.String()); len(stopChar) > 0 {
			return fmt.Errorf(`%w: set '%s'`, ErrArgumentStopCharacter, stopChar)
		}
	}
	if len(a.allowedValues) > 0 && len(a.defaultValues) > 0 {
		for _, valDefault := range a.defaultValues {
			if stopChar := stopCharsPattern.FindString(valDefault.String()); len(stopChar) > 0 {
				return fmt.Errorf(`%w: set '%s'`, ErrArgumentStopCharacter, stopChar)
			}

			contain := false
			for _, valAllowed := range a.allowedValues {
				if valDefault == valAllowed {
					contain = true
					break
				}
			}

			if !contain {
				return fmt.Errorf("%w: set allowed [%s]; set default [%s]",
					ErrArgumentDefaultValueIsNotAllowed,
					strings.Join(coty.ToSliceStrings(a.allowedValues), " "),
					strings.Join(coty.ToSliceStrings(a.defaultValues), " "))
			}
		}
	}
	return nil
}

// ArgumentOpt - source for construct a description of mainName line flag arguments
type ArgumentOpt struct {
	HelpName string `yaml:"help_name"`

	// optional
	IsList        bool     `yaml:"is_list"`
	DefaultValues []string `yaml:"default_values"`
	AllowedValues []string `yaml:"allowed_values"`
	IsOptional    bool     `yaml:"is_optional"`
}

// NewArgument converts opt to AppHelp pointer
func NewArgument(opt *ArgumentOpt) *Argument {
	if opt == nil {
		return nil
	}

	return &Argument{
		helpName: coty.NameArgHelp(opt.HelpName),

		isList:        opt.IsList,
		defaultValues: coty.ToSliceTypesSorted[coty.ArgValue](opt.DefaultValues),
		allowedValues: coty.ToSliceTypesSorted[coty.ArgValue](opt.AllowedValues),
		isOptional:    opt.IsOptional,
	}
}
