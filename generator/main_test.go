package main

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	parsed "github.com/terryhay/dolly/argparser/parsed_data"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/generator/os_decorator"
	osd "github.com/terryhay/dolly/generator/os_decorator"
	"github.com/terryhay/dolly/generator/parser"
	"github.com/terryhay/dolly/utils/dollyerr"
	"os"
	"testing"
)

func TestLogic(t *testing.T) {
	parsingErr := dollyerr.NewError(dollyerr.CodeUndefinedError, fmt.Errorf(gofakeit.Name()))
	configPath := parsed.ArgValue(gofakeit.Name())

	getYAMLConfigErr := dollyerr.NewError(dollyerr.CodeConfigFlagIsNotUsedInCommands, fmt.Errorf(gofakeit.Name()))

	testCases := []struct {
		caseName string

		dollyParseFunc    func(args []string) (res *parsed.ParsedData, err *dollyerr.Error)
		getYAMLConfigFunc func(configPath string) (*confYML.Config, *dollyerr.Error)
		osd               os_decorator.OSDecorator

		expectedErrCode dollyerr.Code
	}{
		{
			caseName: "parsing_error",

			dollyParseFunc: func(arg []string) (res *parsed.ParsedData, err *dollyerr.Error) {
				return nil, parsingErr
			},
			osd: osd.NewOSDecorator(
				&osd.Mock{
					FuncGetArgs: func() []string {
						return nil
					},
				},
			),
			expectedErrCode: parsingErr.Code(),
		},
		{
			caseName: "get_config_path_arg_error",

			dollyParseFunc: func(arg []string) (res *parsed.ParsedData, err *dollyerr.Error) {
				return nil, nil
			},
			osd: osd.NewOSDecorator(
				&osd.Mock{
					FuncGetArgs: func() []string {
						return nil
					},
				},
			),
			expectedErrCode: dollyerr.CodeGeneratorNoRequiredFlag,
		},
		{
			caseName: "get_generate_dir_path_arg_error",

			dollyParseFunc: func(arg []string) (res *parsed.ParsedData, err *dollyerr.Error) {
				return &parsed.ParsedData{
						FlagDataMap: map[apConf.Flag]*parsed.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed.ParsedArgData{
									ArgValues: []parsed.ArgValue{
										configPath,
									},
								},
							},
						},
					},
					nil
			},
			osd: osd.NewOSDecorator(
				&osd.Mock{
					FuncGetArgs: func() []string {
						return nil
					},
				},
			),
			expectedErrCode: dollyerr.CodeGeneratorNoRequiredFlag,
		},
		{
			caseName: "get_yaml_config_error",

			dollyParseFunc: func(arg []string) (res *parsed.ParsedData, err *dollyerr.Error) {
				return &parsed.ParsedData{
						FlagDataMap: map[apConf.Flag]*parsed.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed.ParsedArgData{
									ArgValues: []parsed.ArgValue{
										configPath,
									},
								},
							},
							parser.FlagO: {
								ArgData: &parsed.ParsedArgData{
									ArgValues: []parsed.ArgValue{
										parsed.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*confYML.Config, *dollyerr.Error) {
				return nil, getYAMLConfigErr
			},
			osd: osd.NewOSDecorator(
				&osd.Mock{
					FuncGetArgs: func() []string {
						return nil
					},
				},
			),
			expectedErrCode: dollyerr.CodeConfigFlagIsNotUsedInCommands,
		},
		{
			caseName: "extract_flag_descriptions_error",

			dollyParseFunc: func(arg []string) (res *parsed.ParsedData, err *dollyerr.Error) {
				return &parsed.ParsedData{
						FlagDataMap: map[apConf.Flag]*parsed.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed.ParsedArgData{
									ArgValues: []parsed.ArgValue{
										configPath,
									},
								},
							},
							parser.FlagO: {
								ArgData: &parsed.ParsedArgData{
									ArgValues: []parsed.ArgValue{
										parsed.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*confYML.Config, *dollyerr.Error) {
				return confYML.ConfigSrc{
						ArgParserConfig: confYML.ArgParserConfigSrc{
							FlagDescriptions: []*confYML.FlagDescription{
								nil,
							},
						}.ToConstPtr(),
					}.ToConstPtr(),
					nil
			},
			osd: osd.NewOSDecorator(
				&osd.Mock{
					FuncGetArgs: func() []string {
						return nil
					},
				},
			),
			expectedErrCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "extract_command_descriptions_error",

			dollyParseFunc: func(arg []string) (res *parsed.ParsedData, err *dollyerr.Error) {
				return &parsed.ParsedData{
						FlagDataMap: map[apConf.Flag]*parsed.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed.ParsedArgData{
									ArgValues: []parsed.ArgValue{
										configPath,
									},
								},
							},
							parser.FlagO: {
								ArgData: &parsed.ParsedArgData{
									ArgValues: []parsed.ArgValue{
										parsed.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*confYML.Config, *dollyerr.Error) {
				return confYML.ConfigSrc{
						ArgParserConfig: confYML.ArgParserConfigSrc{
							CommandDescriptions: []*confYML.CommandDescription{
								nil,
							},
						}.ToConstPtr(),
					}.ToConstPtr(),
					nil
			},
			osd: osd.NewOSDecorator(
				&osd.Mock{
					FuncGetArgs: func() []string {
						return nil
					},
				},
			),
			expectedErrCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "checking_error",

			dollyParseFunc: func(arg []string) (res *parsed.ParsedData, err *dollyerr.Error) {
				return &parsed.ParsedData{
						FlagDataMap: map[apConf.Flag]*parsed.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed.ParsedArgData{
									ArgValues: []parsed.ArgValue{
										configPath,
									},
								},
							},
							parser.FlagO: {
								ArgData: &parsed.ParsedArgData{
									ArgValues: []parsed.ArgValue{
										parsed.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*confYML.Config, *dollyerr.Error) {
				return confYML.ConfigSrc{
						ArgParserConfig: confYML.ArgParserConfigSrc{
							CommandDescriptions: []*confYML.CommandDescription{
								confYML.CommandDescriptionSrc{
									Command: gofakeit.Name(),
									RequiredFlags: []string{
										gofakeit.Color(),
									},
								}.ToConstPtr(),
							},
						}.ToConstPtr(),
					}.ToConstPtr(),
					nil
			},
			osd: osd.NewOSDecorator(
				&osd.Mock{
					FuncGetArgs: func() []string {
						return nil
					},
				},
			),
			expectedErrCode: dollyerr.CodeConfigFlagMustHaveDashInFront,
		},
		{
			caseName: "file_write_error",

			dollyParseFunc: func(arg []string) (res *parsed.ParsedData, err *dollyerr.Error) {
				return &parsed.ParsedData{
						FlagDataMap: map[apConf.Flag]*parsed.ParsedFlagData{
							parser.FlagC: {
								ArgData: &parsed.ParsedArgData{
									ArgValues: []parsed.ArgValue{
										configPath,
									},
								},
							},
							parser.FlagO: {
								ArgData: &parsed.ParsedArgData{
									ArgValues: []parsed.ArgValue{
										parsed.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*confYML.Config, *dollyerr.Error) {
				return &confYML.Config{},
					nil
			},
			osd: osd.NewOSDecorator(&osd.Mock{
				FuncGetArgs: func() []string {
					return nil
				},
				FuncIsExist: func(string) bool {
					return false
				},
			}),
			expectedErrCode: dollyerr.CodeGeneratorInvalidPath,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			err, code := logic(tc.dollyParseFunc, tc.getYAMLConfigFunc, tc.osd)
			require.Equal(t, uint(tc.expectedErrCode), code)
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
