package config_yaml

import (
	"fmt"
	"unsafe"
)

// HelpOutTool - type of help out tool
type HelpOutTool uint

const (
	// HelpOutToolPlainText - output help info like plain text
	HelpOutToolPlainText HelpOutTool = iota

	// HelpOutToolManStyle - output help info in a man style stream
	HelpOutToolManStyle
)

// HelpOutConfig - configuration of using help info output tool
type HelpOutConfig struct {
	usingTool HelpOutTool
}

// GetUsingTool - usingTool field getter
func (hoc *HelpOutConfig) GetUsingTool() HelpOutTool {
	if hoc == nil {
		return HelpOutToolPlainText
	}
	return hoc.usingTool
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (hoc *HelpOutConfig) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = hoc

	proxy := struct {
		UsingTool string `yaml:"using_tool"`
	}{}
	if err = unmarshal(&proxy); err != nil {
		return err
	}

	const (
		helpOutToolPlainTextValue = "plain"
		helpOutToolManStyleValue  = "man_style"
	)

	switch proxy.UsingTool {
	case helpOutToolPlainTextValue:
		hoc.usingTool = HelpOutToolPlainText

	case helpOutToolManStyleValue:
		hoc.usingTool = HelpOutToolManStyle

	default:
		return fmt.Errorf(`help_out_config.using_tool unmarshal error: unexpected value %s (expected: "%s", "%s")`,
			proxy.UsingTool, helpOutToolPlainTextValue, helpOutToolManStyleValue)
	}

	return nil
}

// HelpOutConfigSrc - source for constuct a configuration of using help info output tool
type HelpOutConfigSrc struct {
	UsingTool HelpOutTool
}

// ToConstPtr converts src to HelpOutConfig pointer
func (m HelpOutConfigSrc) ToConstPtr() *HelpOutConfig {
	return (*HelpOutConfig)(unsafe.Pointer(&m))
}
