package main

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/internal/generator/config_yaml"
	"github.com/terryhay/dolly/internal/generator/parser"
	"github.com/terryhay/dolly/internal/os_decorator"
	osdMock "github.com/terryhay/dolly/internal/os_decorator/os_decorator_mock"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/parsed_data"
	"os"
	"testing"
)

func TestLogic(t *testing.T) {
	parsingErr := dollyerr.NewError(dollyerr.CodeUndefinedError, fmt.Errorf(gofakeit.Name()))
	configPath := parsed_data.ArgValue(gofakeit.Name())

	getYAMLConfigErr := dollyerr.NewError(dollyerr.CodeConfigFlagIsNotUsedInCommands, fmt.Errorf(gofakeit.Name()))

	testData := []struct {
		caseName string

		dollyParseFunc    func(args []string) (res *parsed_data.ParsedData, err *dollyerr.Error)
		getYAMLConfigFunc func(configPath string) (*config_yaml.Config, *dollyerr.Error)
		osd               os_decorator.OSDecorator

		expectedErrCode dollyerr.Code
	}{
		{
			caseName: "parsing_error",

			dollyParseFunc: func(arg []string) (res *parsed_data.ParsedData, err *dollyerr.Error) {
				return nil, parsingErr
			},
			osd: osdMock.NewOSDecoratorMock(
				osdMock.OSDecoratorMockInit{
					Args: []string{},
				},
			),
			expectedErrCode: parsingErr.Code(),
		},
		{
			caseName: "get_config_path_arg_error",

			dollyParseFunc: func(arg []string) (res *parsed_data.ParsedData, err *dollyerr.Error) {
				return nil, nil
			},
			osd: osdMock.NewOSDecoratorMock(
				osdMock.OSDecoratorMockInit{
					Args: []string{},
				},
			),
			expectedErrCode: dollyerr.CodeGeneratorNoRequiredFlag,
		},
		{
			caseName: "get_generate_dir_path_arg_error",

			dollyParseFunc: func(arg []string) (res *parsed_data.ParsedData, err *dollyerr.Error) {
				return &parsed_data.ParsedData{
						FlagDataMap: map[dollyconf.Flag]*parsed_data.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed_data.ParsedArgData{
									ArgValues: []parsed_data.ArgValue{
										configPath,
									},
								},
							},
						},
					},
					nil
			},
			osd: osdMock.NewOSDecoratorMock(
				osdMock.OSDecoratorMockInit{
					Args: []string{},
				},
			),
			expectedErrCode: dollyerr.CodeGeneratorNoRequiredFlag,
		},
		{
			caseName: "get_yaml_config_error",

			dollyParseFunc: func(arg []string) (res *parsed_data.ParsedData, err *dollyerr.Error) {
				return &parsed_data.ParsedData{
						FlagDataMap: map[dollyconf.Flag]*parsed_data.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed_data.ParsedArgData{
									ArgValues: []parsed_data.ArgValue{
										configPath,
									},
								},
							},
							parser.FlagO: {
								ArgData: &parsed_data.ParsedArgData{
									ArgValues: []parsed_data.ArgValue{
										parsed_data.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*config_yaml.Config, *dollyerr.Error) {
				return nil, getYAMLConfigErr
			},
			osd: osdMock.NewOSDecoratorMock(
				osdMock.OSDecoratorMockInit{
					Args: []string{},
				},
			),
			expectedErrCode: dollyerr.CodeConfigFlagIsNotUsedInCommands,
		},
		{
			caseName: "extract_flag_descriptions_error",

			dollyParseFunc: func(arg []string) (res *parsed_data.ParsedData, err *dollyerr.Error) {
				return &parsed_data.ParsedData{
						FlagDataMap: map[dollyconf.Flag]*parsed_data.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed_data.ParsedArgData{
									ArgValues: []parsed_data.ArgValue{
										configPath,
									},
								},
							},
							parser.FlagO: {
								ArgData: &parsed_data.ParsedArgData{
									ArgValues: []parsed_data.ArgValue{
										parsed_data.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*config_yaml.Config, *dollyerr.Error) {
				return &config_yaml.Config{
						FlagDescriptions: []*config_yaml.FlagDescription{
							nil,
						},
					},
					nil
			},
			osd: osdMock.NewOSDecoratorMock(
				osdMock.OSDecoratorMockInit{
					Args: []string{},
				},
			),
			expectedErrCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "extract_command_descriptions_error",

			dollyParseFunc: func(arg []string) (res *parsed_data.ParsedData, err *dollyerr.Error) {
				return &parsed_data.ParsedData{
						FlagDataMap: map[dollyconf.Flag]*parsed_data.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed_data.ParsedArgData{
									ArgValues: []parsed_data.ArgValue{
										configPath,
									},
								},
							},
							parser.FlagO: {
								ArgData: &parsed_data.ParsedArgData{
									ArgValues: []parsed_data.ArgValue{
										parsed_data.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*config_yaml.Config, *dollyerr.Error) {
				return &config_yaml.Config{
						CommandDescriptions: []*config_yaml.CommandDescription{
							nil,
						},
					},
					nil
			},
			osd: osdMock.NewOSDecoratorMock(
				osdMock.OSDecoratorMockInit{
					Args: []string{},
				},
			),
			expectedErrCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "checking_error",

			dollyParseFunc: func(arg []string) (res *parsed_data.ParsedData, err *dollyerr.Error) {
				return &parsed_data.ParsedData{
						FlagDataMap: map[dollyconf.Flag]*parsed_data.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed_data.ParsedArgData{
									ArgValues: []parsed_data.ArgValue{
										configPath,
									},
								},
							},
							parser.FlagO: {
								ArgData: &parsed_data.ParsedArgData{
									ArgValues: []parsed_data.ArgValue{
										parsed_data.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*config_yaml.Config, *dollyerr.Error) {
				return &config_yaml.Config{
						CommandDescriptions: []*config_yaml.CommandDescription{
							{
								Command: gofakeit.Name(),
								RequiredFlags: []string{
									gofakeit.Color(),
								},
							},
						},
					},
					nil
			},
			osd: osdMock.NewOSDecoratorMock(
				osdMock.OSDecoratorMockInit{
					Args: []string{},
				},
			),
			expectedErrCode: dollyerr.CodeConfigFlagMustHaveDashInFront,
		},
		{
			caseName: "file_write_error",

			dollyParseFunc: func(arg []string) (res *parsed_data.ParsedData, err *dollyerr.Error) {
				return &parsed_data.ParsedData{
						FlagDataMap: map[dollyconf.Flag]*parsed_data.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed_data.ParsedArgData{
									ArgValues: []parsed_data.ArgValue{
										configPath,
									},
								},
							},
							parser.FlagO: {
								ArgData: &parsed_data.ParsedArgData{
									ArgValues: []parsed_data.ArgValue{
										parsed_data.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*config_yaml.Config, *dollyerr.Error) {
				return &config_yaml.Config{},
					nil
			},
			osd: osdMock.NewOSDecoratorMock(osdMock.OSDecoratorMockInit{
				Args: []string{},
				IsNotExistFunc: func(err error) bool {
					return err != nil
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					return nil, dollyerr.NewError(dollyerr.CodeUndefinedError, fmt.Errorf(gofakeit.Name()))
				},
			}),
			expectedErrCode: dollyerr.CodeGeneratorInvalidGeneratePath,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err, code := logic(td.dollyParseFunc, td.getYAMLConfigFunc, td.osd)
			require.Equal(t, uint(td.expectedErrCode), code)
			if code == 0 {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
		})
	}
}

func TestCrasher(t *testing.T) {
	fakeExit := func(int) {
		panic("os.Exit called")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	assert.PanicsWithValue(
		t,
		"os.Exit called",
		func() {
			main()
		},
		"os.Exit was not called")
}
