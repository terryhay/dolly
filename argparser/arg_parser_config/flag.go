package arg_parser_config

import (
	"unsafe"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// Flag contains specification of flag with flag's arguments
type Flag struct {
	nameMain        coty.NameFlag
	namesAdditional map[coty.NameFlag]struct{}
	helpInfo        string
}

// FlagOpt contains source data for cast to Flag
type FlagOpt struct {
	NameMain        coty.NameFlag
	NamesAdditional map[coty.NameFlag]struct{}
	HelpInfo        string
}

// MakeFlag converts opt to Flag pointer
func MakeFlag(opt FlagOpt) *Flag {
	return (*Flag)(unsafe.Pointer(&opt))
}

// GetNameMain gets nameMain field
func (fd *Flag) GetNameMain() coty.NameFlag {
	if fd == nil {
		return coty.NameFlagUndefined
	}
	return fd.nameMain
}

// GetNamesAdditional gets namesAdditional field
func (fd *Flag) GetNamesAdditional() map[coty.NameFlag]struct{} {
	if fd == nil {
		return nil
	}
	return fd.namesAdditional
}

// GetDescriptionHelpInfo gets HelpInfo field
func (fd *Flag) GetDescriptionHelpInfo() string {
	if fd == nil {
		return ""
	}
	return fd.helpInfo
}
