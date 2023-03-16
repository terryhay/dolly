package arg_parser_config

import (
	"sort"
	"strings"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// Command contains specification of command group which contains namesAdditional and arguments
type Command struct {
	nameMain        coty.NameCommand
	namesAdditional map[coty.NameCommand]struct{}
	placeholders    []*Placeholder
	helpInfo        string
}

// CommandOpt contains source data for cast to Command
type CommandOpt struct {
	NameMain        coty.NameCommand
	NamesAdditional map[coty.NameCommand]struct{}
	Placeholders    []*PlaceholderOpt
	HelpInfo        string
}

// NewCommand converts opt to Command pointer
func NewCommand(opt CommandOpt) *Command {
	return &Command{
		nameMain:        opt.NameMain,
		namesAdditional: opt.NamesAdditional,
		placeholders:    createPlaceholders(opt.Placeholders),
		helpInfo:        opt.HelpInfo,
	}
}

// GetNameMain gets nameMain field
func (i *Command) GetNameMain() coty.NameCommand {
	if i == nil {
		return coty.NameCommandUndefined
	}
	return i.nameMain
}

// GetNamesAdditional gets namesAll field
func (i *Command) GetNamesAdditional() map[coty.NameCommand]struct{} {
	if i == nil {
		return nil
	}
	return i.namesAdditional
}

// GetPlaceholders gets placeholders field
func (i *Command) GetPlaceholders() []*Placeholder {
	if i == nil {
		return nil
	}
	return i.placeholders
}

// GetDescriptionHelpInfo gets helpInfo field
func (i *Command) GetDescriptionHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.helpInfo
}

// CreateStringWithCommandNames creates sequence of command names as a string
func (i *Command) CreateStringWithCommandNames() string {
	if i == nil {
		return ""
	}

	names := make([]string, 0, 1+len(i.namesAdditional))
	names = append(names, "")
	for name := range i.namesAdditional {
		names = append(names, name.String())
	}

	sort.Strings(names)
	names[0] = i.nameMain.String()

	return strings.Join(names, ", ")
}

func createPlaceholders(opts []*PlaceholderOpt) []*Placeholder {
	if len(opts) == 0 {
		return nil
	}

	res := make([]*Placeholder, 0, len(opts))
	for _, opt := range opts {
		res = append(res, NewPlaceholder(*opt))
	}

	return res
}
