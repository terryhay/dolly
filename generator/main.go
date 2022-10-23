package main

import (
	"fmt"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/generator/config_checker"
	"github.com/terryhay/dolly/generator/config_data_extractor"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/generator/file_writer"
	"github.com/terryhay/dolly/generator/generate"
	"github.com/terryhay/dolly/generator/os_decorator"
	"github.com/terryhay/dolly/generator/parser"
	"github.com/terryhay/dolly/utils/dollyerr"
)

func main() {
	osd := os_decorator.NewOSDecorator(nil)
	osd.Exit(logic(parser.Parse, confYML.GetConfig, osd))
}

func logic(
	dollyParseFunc func(args []string) (res *parsed_data.ParsedData, err *dollyerr.Error),
	getYAMLConfigFunc func(configPath string) (*confYML.Config, *dollyerr.Error),
	osd os_decorator.OSDecorator,
) (error, uint) {

	argData, err := dollyParseFunc(osd.GetArgs())
	if err != nil {
		return err.Error(), err.Code().ToUint()
	}

	var (
		configYAMLFilePath parsed_data.ArgValue
		contain            bool
	)
	configYAMLFilePath, contain = argData.GetFlagArgValue(parser.FlagC)
	if !contain {
		err = dollyerr.NewError(
			dollyerr.CodeGeneratorNoRequiredFlag,
			fmt.Errorf("parser.generator: can't get required flag \"%s\"", parser.FlagC))
		return err.Error(), err.Code().ToUint()
	}

	var generateDirPath parsed_data.ArgValue
	generateDirPath, contain = argData.GetFlagArgValue(parser.FlagO)
	if !contain {
		err = dollyerr.NewError(
			dollyerr.CodeGeneratorNoRequiredFlag,
			fmt.Errorf("parser.generator: can't get required flag \"%s\"", parser.FlagO))
		return err.Error(), err.Code().ToUint()
	}

	var config *confYML.Config
	config, err = getYAMLConfigFunc(string(configYAMLFilePath))
	if err != nil {
		return err.Error(), err.Code().ToUint()
	}

	var flagDescriptions map[string]*confYML.FlagDescription
	flagDescriptions, err = config_data_extractor.ExtractFlagDescriptionMap(config.GetArgParserConfig().GetFlagDescriptions())
	if err != nil {
		return err.Error(), err.Code().ToUint()
	}

	var commandDescriptions map[string]*confYML.CommandDescription
	commandDescriptions, err = config_data_extractor.ExtractCommandDescriptionMap(config.GetArgParserConfig().GetCommandDescriptions())
	if err != nil {
		return err.Error(), err.Code().ToUint()
	}

	err = config_checker.Check(config.GetArgParserConfig().GetNamelessCommandDescription(), commandDescriptions, flagDescriptions)
	if err != nil {
		return err.Error(), err.Code().ToUint()
	}

	err = file_writer.WriteFile(osd, string(generateDirPath), generate.Generate(config.GetArgParserConfig(), config.GetHelpOutConfig(), flagDescriptions))
	return err.Error(), err.Code().ToUint()
}
