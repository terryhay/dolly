package main

import (
	"fmt"
	parsed "github.com/terryhay/dolly/argparser/parsed_data"
	confCheck "github.com/terryhay/dolly/generator/config_checker"
	conf "github.com/terryhay/dolly/generator/config_data_extractor"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	write "github.com/terryhay/dolly/generator/file_writer"
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
	dollyParseFunc func(args []string) (res *parsed.ParsedData, err *dollyerr.Error),
	getYAMLConfigFunc func(configPath string) (*confYML.Config, *dollyerr.Error),
	osd os_decorator.OSDecorator,
) (error, uint) {

	argData, err := dollyParseFunc(osd.GetArgs())
	if err != nil {
		return err.Error(), err.Code().ToUint()
	}

	var (
		configYAMLFilePath parsed.ArgValue
		contain            bool
	)
	configYAMLFilePath, contain = argData.GetFlagArgValue(parser.FlagC)
	if !contain {
		err = dollyerr.NewError(
			dollyerr.CodeGeneratorNoRequiredFlag,
			fmt.Errorf("parser.generator: can't get required flag \"%s\"", parser.FlagC))
		return err.Error(), err.Code().ToUint()
	}

	var generateDirPath parsed.ArgValue
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
	if flagDescriptions, err = conf.ExtractFlagDescriptionMap(
		config.GetArgParserConfig().GetFlagDescriptions(),
	); err != nil {
		return err.Error(), err.Code().ToUint()
	}

	var commandDescriptions map[string]*confYML.CommandDescription
	if commandDescriptions, err = conf.ExtractCommandDescriptionMap(
		config.GetArgParserConfig().GetCommandDescriptions(),
	); err != nil {
		return err.Error(), err.Code().ToUint()
	}

	if err = confCheck.Check(
		config.GetArgParserConfig().GetNamelessCommandDescription(), commandDescriptions, flagDescriptions,
	); err != nil {
		return err.Error(), err.Code().ToUint()
	}

	err = write.WriteFile(
		osd,
		generateDirPath.ToString(),
		generate.Generate(config.GetArgParserConfig(), config.GetHelpOutConfig(), flagDescriptions),
	)
	return err.Error(), err.Code().ToUint()
}
