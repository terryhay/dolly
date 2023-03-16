package config_yaml

import (
	"errors"
	"fmt"
)

// HelpOutTool - type of help out tool
type HelpOutTool uint

const (
	// HelpOutToolUndefined - undefined value of HelpOutTool
	HelpOutToolUndefined HelpOutTool = iota

	// HelpOutToolPlainText - output help info like plain text
	HelpOutToolPlainText

	// HelpOutToolManStyle - output help info in a man style stream
	HelpOutToolManStyle
)

func makeHelpOutTool(v string) HelpOutTool {
	switch {
	case len(v) == 0, v == usingToolPlain:
		return HelpOutToolPlainText

	case v == usingToolManStyle:
		return HelpOutToolManStyle

	default:
		return HelpOutToolUndefined
	}
}

const (
	usingToolPlain    = "plain"
	usingToolManStyle = "man_style"
)

// HelpOutConfig - configuration of using help info output tool
type HelpOutConfig struct {
	usingTool HelpOutTool
}

// GetUsingTool gets usingTool field
func (hoc *HelpOutConfig) GetUsingTool() HelpOutTool {
	if hoc == nil {
		return HelpOutToolPlainText
	}
	return hoc.usingTool
}

// ErrHelpOutUnexpectedTool - 'help_out_config.using_tool' is empty or invalid in source yaml file
var ErrHelpOutUnexpectedTool = errors.New(`HelpOutConfig.IsValid: 'help_out_config.using_tool' is empty or invalid in source yaml file`)

// IsValid checks if HelpOutConfig is valid
func (hoc *HelpOutConfig) IsValid() error {
	if hoc == nil {
		return nil
	}

	if hoc.usingTool == HelpOutToolUndefined {
		return fmt.Errorf(`%w: expected: [%s, %s]`, ErrHelpOutUnexpectedTool, usingToolPlain, usingToolManStyle)
	}

	return nil
}

// HelpOutConfigOpt - source for construct a configuration of using help info output tool
type HelpOutConfigOpt struct {
	UsingTool string `yaml:"using_tool"`
}

// NewHelpOutConfig converts opt to HelpOutConfig pointer
func NewHelpOutConfig(opt *HelpOutConfigOpt) *HelpOutConfig {
	if opt == nil {
		return nil
	}

	return &HelpOutConfig{
		usingTool: makeHelpOutTool(opt.UsingTool),
	}
}
