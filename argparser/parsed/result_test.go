package parsed

import (
	"testing"

	"github.com/stretchr/testify/require"
	clArg "github.com/terryhay/dolly/argparser/command_line_argument"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestParsedDataOpt(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var opt *ResultOpt

		opt.SetFlagName(coty.RandIDPlaceholder(), coty.RandNameFlagShort())
		opt.SetArg(coty.RandIDPlaceholder(), MakeArgValue(clArg.RandCmdArg()))

		require.Nil(t, opt.PlaceholderByID(coty.RandIDPlaceholder()))
		require.True(t, opt.PlaceholderDoesNotHaveArgs(coty.RandIDPlaceholder()))
		require.Nil(t, MakeResult(opt))
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		t.Parallel()

		{
			pointer := &ResultOpt{
				CommandMainName: coty.RandNameCommand(),
			}

			pointer.SetFlagName(coty.RandIDPlaceholder(), coty.RandNameFlagShort())
			require.NotNil(t, pointer.PlaceholderByID(coty.RandIDPlaceholder()))
			require.True(t, pointer.PlaceholderDoesNotHaveArgs(coty.RandIDPlaceholder()))
		}
		{
			pointer := &ResultOpt{
				CommandMainName: coty.RandNameCommand(),
			}
			require.True(t, pointer.PlaceholderDoesNotHaveArgs(coty.RandIDPlaceholder()))

			pointer.SetArg(coty.RandIDPlaceholder(), MakeArgValue(clArg.RandCmdArg()))
			require.NotNil(t, pointer.PlaceholderByID(coty.RandIDPlaceholder()))
			require.False(t, pointer.PlaceholderDoesNotHaveArgs(coty.RandIDPlaceholder()))
		}
	})
}

func TestParsedData(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *Result

		vals, err := pointer.FlagArgValues(coty.RandNameFlagShort())
		require.Nil(t, vals)
		require.ErrorIs(t, err, ErrFlagArgValuesNilPointer)

		require.Equal(t, coty.NameCommandUndefined, pointer.GetCommandMainName())
		require.Nil(t, pointer.GetPlaceholdersByIDs())
		require.Nil(t, pointer.PlaceholderByID(coty.RandIDPlaceholder()))

		var val ArgValue
		val, err = pointer.FlagArgValue(coty.RandNameFlagShort())
		require.Equal(t, ArgValueDefault, val)
		require.NotNil(t, err)
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		t.Parallel()

		opt := ResultOpt{
			CommandMainName: coty.RandNameCommand(),
		}

		pointer := MakeResult(&opt)
		require.Equal(t, opt.CommandMainName, pointer.GetCommandMainName())
		require.Equal(t, 0, len(pointer.GetPlaceholdersByIDs()))

		opt.SetFlagName(coty.RandIDPlaceholder(), coty.RandNameFlagShort())
		opt.SetArg(coty.RandIDPlaceholder(), RandArgValue())
		opt.SetArg(coty.RandIDPlaceholder(), RandArgValueSecond())
		pointer = MakeResult(&opt)

		expPlaceholder := NewPlaceholder(&PlaceholderOpt{
			ID:   coty.RandIDPlaceholder(),
			Flag: coty.RandNameFlagShort(),
			Argument: &ArgumentOpt{
				ArgValues: []ArgValue{
					RandArgValue(),
					RandArgValueSecond(),
				},
			},
		})

		require.Equal(t, expPlaceholder, pointer.PlaceholderByID(coty.RandIDPlaceholder()))

		require.Equal(t, map[coty.IDPlaceholder]*Placeholder{
			coty.RandIDPlaceholder(): expPlaceholder,
		}, pointer.GetPlaceholdersByIDs())

		_, err := pointer.FlagArgValues(coty.RandNameFlagShort())
		require.NoError(t, err)
	})
}

func TestGetFlagArgValuesErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string

		parsedData *Result
		flagName   coty.NameFlag

		expError error
	}{
		{
			caseName: "nil_pointer",
			expError: ErrFlagArgValuesNilPointer,
		},
		{
			caseName:   "no_flag_data",
			parsedData: &Result{},
			expError:   ErrFlagArgValuesNoPlaceholder,
		},
		{
			caseName: "no_flag_arg",
			flagName: coty.RandNameFlagShort(),
			parsedData: MakeResult(&ResultOpt{
				PlaceholdersByID: map[coty.IDPlaceholder]*PlaceholderOpt{
					coty.RandIDPlaceholder(): {
						ID:   coty.RandIDPlaceholder(),
						Flag: coty.RandNameFlagShort(),
					},
				},
			}),
			expError: ErrFlagArgValuesNoArgValues,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName+"_GetFlagArgValues", func(t *testing.T) {
			t.Parallel()

			v, err := tc.parsedData.FlagArgValues(tc.flagName)
			if tc.expError == nil {
				require.Equal(t, 1, len(v))
				require.Nil(t, err)
				return
			}

			require.Equal(t, 0, len(v))
			require.ErrorIs(t, err, tc.expError)
		})
	}
}

func TestFlagArgValueErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string

		parsedData *Result
		flag       coty.NameFlag

		expError error
	}{
		{
			caseName: "nil_pointer",
			expError: ErrFlagArgValuesNilPointer,
		},
		{
			caseName:   "empty_result",
			parsedData: &Result{},
			expError:   ErrFlagArgValuesNoPlaceholder,
		},
		{
			caseName: "no_flag_data",
			flag:     coty.RandNameFlagShort(),
			parsedData: MakeResult(&ResultOpt{
				PlaceholdersByID: map[coty.IDPlaceholder]*PlaceholderOpt{
					coty.RandIDPlaceholder(): {
						ID:   coty.RandIDPlaceholder(),
						Flag: coty.RandNameFlagShort(),
					},
				},
			}),
			expError: ErrFlagArgValuesNoArgValues,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName+"_GetFlagArgValues", func(t *testing.T) {
			t.Parallel()

			v, err := tc.parsedData.FlagArgValue(tc.flag)
			if tc.expError == nil {
				require.Equal(t, 1, len(v))
				require.Nil(t, err)
				return
			}

			require.Equal(t, 0, len(v))
			require.ErrorIs(t, err, tc.expError)
		})
	}
}

func TestFlagArgValue(t *testing.T) {
	t.Parallel()

	parsedData := MakeResult(&ResultOpt{
		PlaceholdersByID: map[coty.IDPlaceholder]*PlaceholderOpt{
			coty.RandIDPlaceholder(): {
				ID:   coty.RandIDPlaceholder(),
				Flag: coty.RandNameFlagShort(),
				Argument: &ArgumentOpt{
					ArgValues: []ArgValue{
						RandArgValue(),
						RandArgValueSecond(),
					},
				},
			},
			coty.RandIDPlaceholderSecond(): {
				ID:   coty.RandIDPlaceholderSecond(),
				Flag: coty.RandNameFlagLong(),
			},
			coty.RandIDPlaceholderThird(): {
				ID:   coty.RandIDPlaceholderThird(),
				Flag: coty.RandNameFlagOneLetter(),
				Argument: &ArgumentOpt{
					ArgValues: []ArgValue{
						RandArgValueThird(),
					},
				},
			},
		},
	})

	t.Run("FlagArgValues", func(t *testing.T) {
		t.Parallel()

		v, err := parsedData.FlagArgValues(coty.RandNameFlagShort())
		require.NoError(t, err)
		require.Equal(t, RandArgValue(), v[0])

		v, err = parsedData.FlagArgValues(coty.RandNameFlagLong())
		require.ErrorIs(t, err, ErrFlagArgValuesNoArgValues)
		require.Nil(t, v)

		v, err = parsedData.FlagArgValues(coty.RandNameFlagOneLetter())
		require.NoError(t, err)
		require.Equal(t, RandArgValueThird(), v[0])
	})

	t.Run("FlagArgValue", func(t *testing.T) {
		t.Parallel()

		v, err := parsedData.FlagArgValue(coty.RandNameFlagShort())
		require.Nil(t, err)
		require.Equal(t, RandArgValue(), v)
	})
}

func TestPlaceholderArgValueErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string

		parsedData *Result
		id         coty.IDPlaceholder

		expError error
	}{
		{
			caseName: "nil_pointer",
			expError: ErrPlaceholderArgValuesNilPointer,
		},
		{
			caseName:   "empty_result",
			parsedData: &Result{},
			expError:   ErrPlaceholderArgValuesNoPlaceholder,
		},
		{
			caseName: "no_placeholder_data",
			id:       coty.RandIDPlaceholder(),
			parsedData: MakeResult(&ResultOpt{
				PlaceholdersByID: map[coty.IDPlaceholder]*PlaceholderOpt{
					coty.RandIDPlaceholder(): {
						ID:   coty.RandIDPlaceholder(),
						Flag: coty.RandNameFlagShort(),
					},
				},
			}),
			expError: ErrPlaceholderArgValuesNoArgValues,
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseName+"_GetFlagArgValues", func(t *testing.T) {
			v, err := tc.parsedData.PlaceholderArgValue(tc.id)
			if tc.expError == nil {
				require.Equal(t, 1, len(v))
				require.Nil(t, err)
				return
			}

			require.Equal(t, 0, len(v))
			require.ErrorIs(t, err, tc.expError)
		})
	}
}

func TestPlaceholderArgValue(t *testing.T) {
	t.Parallel()

	parsedData := MakeResult(&ResultOpt{
		PlaceholdersByID: map[coty.IDPlaceholder]*PlaceholderOpt{
			coty.RandIDPlaceholder(): {
				ID:   coty.RandIDPlaceholder(),
				Flag: coty.RandNameFlagShort(),
				Argument: &ArgumentOpt{
					ArgValues: []ArgValue{
						RandArgValue(),
						RandArgValueSecond(),
					},
				},
			},
			coty.RandIDPlaceholderSecond(): {
				ID:   coty.RandIDPlaceholderSecond(),
				Flag: coty.RandNameFlagLong(),
			},
			coty.RandIDPlaceholderThird(): {
				ID:   coty.RandIDPlaceholderThird(),
				Flag: coty.RandNameFlagOneLetter(),
				Argument: &ArgumentOpt{
					ArgValues: []ArgValue{
						RandArgValueThird(),
					},
				},
			},
		},
	})

	t.Run("PlaceholderArgValues", func(t *testing.T) {
		v, err := parsedData.PlaceholderArgValues(coty.RandIDPlaceholder())
		require.NoError(t, err)
		require.Equal(t, RandArgValue(), v[0])

		v, err = parsedData.PlaceholderArgValues(coty.RandIDPlaceholderSecond())
		require.ErrorIs(t, err, ErrPlaceholderArgValuesNoArgValues)
		require.Nil(t, v)

		v, err = parsedData.PlaceholderArgValues(coty.RandIDPlaceholderThird())
		require.NoError(t, err)
		require.Equal(t, RandArgValueThird(), v[0])
	})

	t.Run("PlaceholderArgValue", func(t *testing.T) {
		v, err := parsedData.PlaceholderArgValue(coty.RandIDPlaceholder())
		require.Nil(t, err)
		require.Equal(t, RandArgValue(), v)
	})
}
