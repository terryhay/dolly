package main

import (
	"errors"
	"math/rand"
	"os"
	"testing"

	"github.com/terryhay/dolly/generator/proxyes/file_proxy"
	"github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"bou.ke/monkey"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/argparser/parsed"
	ce "github.com/terryhay/dolly/generator/config_entity"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/generator/parser"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestProcess(t *testing.T) {
	t.Parallel()

	errParsing := errors.New(gofakeit.Paragraph(1, 1, rand.Intn(10)+1, ""))
	configPath := parsed.ArgValue(gofakeit.Name())
	errYmlConfig := errors.New(gofakeit.Paragraph(1, 1, rand.Intn(10)+1, ""))
	errExist := errors.New(gofakeit.Paragraph(1, 1, rand.Intn(10)+1, ""))

	tests := []struct {
		caseName string

		funcParse          func(args []string) (res *parsed.Result, err error)
		funcLoadYamlConfig func(decOS os_proxy.Proxy, configPath string) (*confYML.Config, error)
		proxyOS            os_proxy.Proxy

		expErr      error
		expCodeExit os_proxy.ExitCode
	}{
		{
			caseName: "get_arguments_error",

			proxyOS: os_proxy.Mock(os_proxy.Opt{}),

			expErr:      os_proxy.ErrGetArgsNoImplementation,
			expCodeExit: ExitCodeGetArgsError,
		},
		{
			caseName: "parsing_error",

			funcParse: func(arg []string) (res *parsed.Result, err error) {
				return nil, errParsing
			},
			proxyOS: os_proxy.Mock(os_proxy.Opt{
				SlotGetArgs: func() []string {
					return nil
				},
			}),

			expErr:      errParsing,
			expCodeExit: ExitCodeArgParseError,
		},
		{
			caseName: "get_config_path_arg_error",

			funcParse: func(arg []string) (res *parsed.Result, err error) {
				return nil, nil
			},
			proxyOS: os_proxy.Mock(os_proxy.Opt{
				SlotGetArgs: func() []string {
					return nil
				},
			}),

			expErr:      parsed.ErrFlagArgValuesNilPointer,
			expCodeExit: ExitCodeGetFlagArgValueError,
		},
		{
			caseName: "get_generate_dir_path_arg_error",

			funcParse: func(arg []string) (res *parsed.Result, err error) {
				return parsed.MakeResult(&parsed.ResultOpt{
						PlaceholdersByID: map[coty.IDPlaceholder]*parsed.PlaceholderOpt{
							coty.ArgPlaceholderIDUndefined: {
								ID:   coty.ArgPlaceholderIDUndefined,
								Flag: parser.FlagC,
								Argument: &parsed.ArgumentOpt{
									ArgValues: []parsed.ArgValue{
										configPath,
									},
								},
							},
						},
					}),
					nil
			},
			proxyOS: os_proxy.Mock(os_proxy.Opt{
				SlotGetArgs: func() []string {
					return nil
				},
			}),

			expErr:      parsed.ErrFlagArgValuesNoPlaceholder,
			expCodeExit: ExitCodeGetFlagArgValueError,
		},
		{
			caseName: "load_yaml_config_error",

			funcParse: func(arg []string) (res *parsed.Result, err error) {
				return parsed.MakeResult(&parsed.ResultOpt{
						PlaceholdersByID: map[coty.IDPlaceholder]*parsed.PlaceholderOpt{
							coty.RandIDPlaceholder(): {
								ID:   coty.RandIDPlaceholder(),
								Flag: parser.FlagC,
								Argument: &parsed.ArgumentOpt{
									ArgValues: []parsed.ArgValue{
										configPath,
									},
								},
							},

							coty.RandIDPlaceholderSecond(): {
								ID:   coty.RandIDPlaceholderSecond(),
								Flag: parser.FlagO,
								Argument: &parsed.ArgumentOpt{
									ArgValues: []parsed.ArgValue{
										parsed.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					}),
					nil
			},
			funcLoadYamlConfig: func(_ os_proxy.Proxy, configPath string) (*confYML.Config, error) {
				return nil, errYmlConfig
			},
			proxyOS: os_proxy.Mock(os_proxy.Opt{
				SlotGetArgs: func() []string {
					return nil
				},
			}),

			expErr:      errYmlConfig,
			expCodeExit: ExitCodeLoadParseConfigError,
		},
		{
			caseName: "make_config_entity_error",

			funcParse: func(arg []string) (res *parsed.Result, err error) {
				return parsed.MakeResult(&parsed.ResultOpt{
					PlaceholdersByID: map[coty.IDPlaceholder]*parsed.PlaceholderOpt{
						coty.RandIDPlaceholder(): {
							ID:   coty.RandIDPlaceholder(),
							Flag: parser.FlagC,
							Argument: &parsed.ArgumentOpt{
								ArgValues: []parsed.ArgValue{
									configPath,
								},
							},
						},

						coty.RandIDPlaceholderSecond(): {
							ID:   coty.RandIDPlaceholderSecond(),
							Flag: parser.FlagO,
							Argument: &parsed.ArgumentOpt{
								ArgValues: []parsed.ArgValue{
									parsed.ArgValue(gofakeit.Name()),
								},
							},
						},
					},
				}), nil
			},
			funcLoadYamlConfig: func(_ os_proxy.Proxy, configPath string) (*confYML.Config, error) {
				return confYML.NewConfig(&confYML.ConfigOpt{
						Version: "1.0.0",
						ArgParserConfig: &confYML.ArgParserConfigOpt{
							AppHelp: &confYML.AppHelpOpt{
								AppName:         coty.RandNameApp().String(),
								ChapterNameInfo: coty.RandInfoChapterDescription().String(),
							},
							NamelessCommand: &confYML.NamelessCommandOpt{
								ChapterDescriptionInfo: coty.RandInfoChapterDescription().String(),
							},
						},
					}),
					nil
			},
			proxyOS: os_proxy.Mock(os_proxy.Opt{
				SlotGetArgs: func() []string {
					return nil
				},
			}),

			expErr:      ce.ErrMakeConfigEntity,
			expCodeExit: ExitCodeConfigEntityMakeError,
		},
		{
			caseName: "file_write_error",

			funcParse: func(arg []string) (res *parsed.Result, err error) {
				return parsed.MakeResult(&parsed.ResultOpt{
						PlaceholdersByID: map[coty.IDPlaceholder]*parsed.PlaceholderOpt{
							coty.RandIDPlaceholder(): {
								ID:   coty.RandIDPlaceholder(),
								Flag: parser.FlagC,
								Argument: &parsed.ArgumentOpt{
									ArgValues: []parsed.ArgValue{
										configPath,
									},
								},
							},

							coty.RandIDPlaceholderSecond(): {
								ID:   coty.RandIDPlaceholderSecond(),
								Flag: parser.FlagO,
								Argument: &parsed.ArgumentOpt{
									ArgValues: []parsed.ArgValue{
										parsed.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					}),
					nil
			},
			funcLoadYamlConfig: func(_ os_proxy.Proxy, configPath string) (*confYML.Config, error) {
				return confYML.NewConfig(&confYML.ConfigOpt{
						Version: "1.0.0",
						ArgParserConfig: &confYML.ArgParserConfigOpt{
							AppHelp: &confYML.AppHelpOpt{
								AppName:         coty.RandNameApp().String(),
								ChapterNameInfo: coty.RandInfoChapterDescription().String(),
							},
							NamelessCommand: &confYML.NamelessCommandOpt{
								ChapterDescriptionInfo: coty.RandInfoChapterDescription().String(),
							},
							HelpCommand: &confYML.HelpCommandOpt{
								MainName: "-h",
							},
						},
					}),
					nil
			},
			proxyOS: os_proxy.Mock(os_proxy.Opt{
				SlotGetArgs: func() []string { return nil },
				SlotIsExist: func(string) error { return errExist },
			}),

			expErr:      errExist,
			expCodeExit: ExitCodeWriteFileError,
		},
		{
			caseName: "success",
			funcParse: func(arg []string) (res *parsed.Result, err error) {
				return parsed.MakeResult(&parsed.ResultOpt{
						PlaceholdersByID: map[coty.IDPlaceholder]*parsed.PlaceholderOpt{
							coty.RandIDPlaceholder(): {
								ID:   coty.RandIDPlaceholder(),
								Flag: parser.FlagC,
								Argument: &parsed.ArgumentOpt{
									ArgValues: []parsed.ArgValue{
										configPath,
									},
								},
							},

							coty.RandIDPlaceholderSecond(): {
								ID:   coty.RandIDPlaceholderSecond(),
								Flag: parser.FlagO,
								Argument: &parsed.ArgumentOpt{
									ArgValues: []parsed.ArgValue{
										parsed.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					}),
					nil
			},
			funcLoadYamlConfig: func(_ os_proxy.Proxy, configPath string) (*confYML.Config, error) {
				return confYML.NewConfig(&confYML.ConfigOpt{
						Version: "1.0.0",
						ArgParserConfig: &confYML.ArgParserConfigOpt{
							AppHelp: &confYML.AppHelpOpt{
								AppName:         coty.RandNameApp().String(),
								ChapterNameInfo: coty.RandInfoChapterDescription().String(),
							},
							NamelessCommand: &confYML.NamelessCommandOpt{
								ChapterDescriptionInfo: coty.RandInfoChapterDescription().String(),
							},
							HelpCommand: &confYML.HelpCommandOpt{
								MainName: "-h",
							},
						},
					}),
					nil
			},
			proxyOS: os_proxy.Mock(os_proxy.Opt{
				SlotGetArgs: func() []string { return nil },
				SlotCreate: func(path string) (file_proxy.Proxy, error) {
					return file_proxy.Mock(file_proxy.Opt{
						SlotClose:       func() error { return nil },
						SlotWriteString: func(string) error { return nil },
					}), nil
				},
				SlotIsExist: func(string) error { return nil },
			}),

			expCodeExit: ExitCodeSuccess,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			codeExit, err := process(tc.proxyOS, tc.funcParse, tc.funcLoadYamlConfig)
			require.ErrorIs(t, err, tc.expErr)
			require.Equal(t, tc.expCodeExit, codeExit)

			if codeExit != ExitCodeSuccess {
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestCrasher(t *testing.T) {
	t.Parallel()

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
