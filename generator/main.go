package main

import (
	"github.com/terryhay/dolly/argparser/parsed"
	ce "github.com/terryhay/dolly/generator/config_entity"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	write "github.com/terryhay/dolly/generator/file_writer"
	"github.com/terryhay/dolly/generator/generate"
	"github.com/terryhay/dolly/generator/parser"
	"github.com/terryhay/dolly/generator/proxyes/os_proxy"
)

const (
	// ExitCodeSuccess - successful completion of the program
	ExitCodeSuccess os_proxy.ExitCode = iota

	// ExitCodeGetArgsError - os.GetArgs error
	ExitCodeGetArgsError

	// ExitCodeArgParseError - internal error of command line argument parsing
	ExitCodeArgParseError

	// ExitCodeGetFlagArgValueError - can't get flag argument value
	ExitCodeGetFlagArgValueError

	// ExitCodeLoadParseConfigError - can't load parse config
	ExitCodeLoadParseConfigError

	// ExitCodeConfigEntityMakeError - can't create ConfigEntity object
	ExitCodeConfigEntityMakeError

	// ExitCodeWriteFileError - can't write file
	ExitCodeWriteFileError
)

func main() {
	proxyOS := os_proxy.New()
	proxyOS.Exit(process(proxyOS, parser.Parse, confYML.Load))
}

func process(
	proxyOS os_proxy.Proxy,
	funcParse func(args []string) (res *parsed.Result, err error),
	funcLoadParseConfig func(decOS os_proxy.Proxy, configPath string) (*confYML.Config, error),
) (
	os_proxy.ExitCode,
	error,
) {
	args, errArgs := proxyOS.GetArgs()
	if errArgs != nil {
		return ExitCodeGetArgsError, errArgs
	}

	argParsed, errParse := funcParse(args)
	if errParse != nil {
		return ExitCodeArgParseError, errParse
	}

	uploadConfigPath, errFlagC := argParsed.FlagArgValue(parser.FlagC)
	if errFlagC != nil {
		return ExitCodeGetFlagArgValueError, errFlagC
	}

	generateDirPath, errFlagO := argParsed.FlagArgValue(parser.FlagO)
	if errFlagO != nil {
		return ExitCodeGetFlagArgValueError, errFlagO
	}

	config, errLoad := funcLoadParseConfig(proxyOS, string(uploadConfigPath))
	if errLoad != nil {
		return ExitCodeLoadParseConfigError, errLoad
	}

	configEntity, errConfig := ce.MakeConfigEntity(config)
	if errConfig != nil {
		return ExitCodeConfigEntityMakeError, errConfig
	}

	if errWrite := write.WriteFile(proxyOS, generateDirPath.String(), generate.Generate(configEntity)); errWrite != nil {
		return ExitCodeWriteFileError, errWrite
	}

	return ExitCodeSuccess, nil
}
