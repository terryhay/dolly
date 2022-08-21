package main

import (
	"fmt"
	"github.com/terryhay/dolly/internal/generator/config_checker"
	"github.com/terryhay/dolly/internal/generator/config_data_extractor"
	"github.com/terryhay/dolly/internal/generator/config_yaml"
	"github.com/terryhay/dolly/internal/generator/file_writer"
	"github.com/terryhay/dolly/internal/generator/generate"
	"github.com/terryhay/dolly/internal/generator/parser"
	"github.com/terryhay/dolly/internal/os_decorator"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/parsed_data"
)

func main() {
	osd := os_decorator.NewOSDecorator()
	osd.Exit(logic(parser.Parse, config_yaml.GetConfig, osd))
}

func logic(
	dollyParseFunc func(args []string) (res *parsed_data.ParsedData, err *dollyerr.Error),
	getYAMLConfigFunc func(configPath string) (*config_yaml.Config, *dollyerr.Error),
	osd os_decorator.OSDecorator,
) (error, uint) {

	argData, err := dollyParseFunc(osd.GetArgs())
	if err != nil {
		return err, err.Code().ToUint()
	}

	var (
		configYAMLFilePath parsed_data.ArgValue
		contain            bool
	)
	configYAMLFilePath, contain = argData.GetFlagArgValue(parser.FlagC)
	if !contain {
		err = dollyerr.NewError(
			dollyerr.CodeGeneratorNoRequiredFlag,
			fmt.Errorf("parser.generator: can't get required flag \"%v\"", parser.FlagC))
		return err, err.Code().ToUint()
	}

	var generateDirPath parsed_data.ArgValue
	generateDirPath, contain = argData.GetFlagArgValue(parser.FlagO)
	if !contain {
		err = dollyerr.NewError(
			dollyerr.CodeGeneratorNoRequiredFlag,
			fmt.Errorf("parser.generator: can't get required flag \"%v\"", parser.FlagO))
		return err, err.Code().ToUint()
	}

	var config *config_yaml.Config
	config, err = getYAMLConfigFunc(string(configYAMLFilePath))
	if err != nil {
		return err, err.Code().ToUint()
	}

	var flagDescriptions map[string]*config_yaml.FlagDescription
	flagDescriptions, err = config_data_extractor.ExtractFlagDescriptionMap(config.GetFlagDescriptions())
	if err != nil {
		return err, err.Code().ToUint()
	}

	var commandDescriptions map[string]*config_yaml.CommandDescription
	commandDescriptions, err = config_data_extractor.ExtractCommandDescriptionMap(config.GetCommandDescriptions())
	if err != nil {
		return err, err.Code().ToUint()
	}

	err = config_checker.Check(config.GetNamelessCommandDescription(), commandDescriptions, flagDescriptions)
	if err != nil {
		return err, err.Code().ToUint()
	}

	err = file_writer.Write(osd, string(generateDirPath), generate.Generate(config, flagDescriptions))
	return err, err.Code().ToUint()
}
