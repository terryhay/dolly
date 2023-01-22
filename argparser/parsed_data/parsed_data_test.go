package parsed_data

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"testing"
)

func TestParsedDataGetters(t *testing.T) {
	t.Parallel()

	var pointer *ParsedData

	{
		require.Equal(t, apConf.CommandIDUndefined, pointer.GetCommandID())
		require.Equal(t, apConf.CommandUndefined, pointer.GetCommand())
		require.Nil(t, pointer.GetAgrData())
		require.Nil(t, pointer.GetFlagDataMap())
	}
	{
		pointer = NewParsedData(
			apConf.CommandID(gofakeit.Uint32()),
			apConf.Command(gofakeit.Name()),
			NewParsedArgData(nil),
			map[apConf.Flag]*ParsedFlagData{
				apConf.Flag(gofakeit.Name()): NewParsedFlagData(
					apConf.Flag(gofakeit.Name()),
					NewParsedArgData(nil)),
			},
		)

		require.Equal(t, pointer.CommandID, pointer.GetCommandID())
		require.Equal(t, pointer.Command, pointer.GetCommand())
		require.Equal(t, pointer.ArgData, pointer.GetAgrData())
		require.Equal(t, pointer.FlagDataMap, pointer.GetFlagDataMap())
	}
	{
		pointer = NewParsedData(
			apConf.CommandID(gofakeit.Uint32()),
			apConf.Command(gofakeit.Name()),
			NewParsedArgData(nil),
			nil,
		)

		require.Nil(t, pointer.GetFlagDataMap())
	}
}

func TestGetFlagArgValuesErrors(t *testing.T) {
	t.Parallel()

	flag := apConf.Flag(gofakeit.Color())

	testCases := []struct {
		caseName string

		parsedData *ParsedData
		flag       apConf.Flag

		expectedSuccess bool
	}{
		{
			caseName: "nil_pointer",
		},
		{
			caseName:   "no_flag_data",
			parsedData: &ParsedData{},
		},
		{
			caseName: "no_flag_data",
			flag:     flag,
			parsedData: &ParsedData{
				FlagDataMap: map[apConf.Flag]*ParsedFlagData{
					flag: {
						Flag:    flag,
						ArgData: &ParsedArgData{},
					},
				},
			},
			expectedSuccess: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName+"_GetFlagArgValues", func(t *testing.T) {
			v, ok := tc.parsedData.GetFlagArgValues(tc.flag)
			require.Equal(t, 0, len(v))
			require.Equal(t, tc.expectedSuccess, ok)
		})
	}
}

func TestGetFlagArgValueErrors(t *testing.T) {
	t.Parallel()

	flag := apConf.Flag(gofakeit.Color())

	testCases := []struct {
		caseName string

		parsedData *ParsedData
		flag       apConf.Flag

		expectedSuccess bool
	}{
		{
			caseName: "nil_pointer",
		},
		{
			caseName:   "no_flag_data",
			parsedData: &ParsedData{},
		},
		{
			caseName: "no_flag_data",
			flag:     flag,
			parsedData: &ParsedData{
				FlagDataMap: map[apConf.Flag]*ParsedFlagData{
					flag: {
						Flag:    flag,
						ArgData: &ParsedArgData{},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName+"_GetFlagArgValues", func(t *testing.T) {
			v, ok := tc.parsedData.GetFlagArgValue(tc.flag)
			require.Equal(t, 0, len(v))
			require.Equal(t, tc.expectedSuccess, ok)
		})
	}
}

func TestGetFlagArgValue(t *testing.T) {
	t.Parallel()

	flag := apConf.Flag(gofakeit.Color())
	value := ArgValue(gofakeit.Color())

	parsedData := &ParsedData{
		FlagDataMap: map[apConf.Flag]*ParsedFlagData{
			flag: {
				Flag: flag,
				ArgData: &ParsedArgData{
					ArgValues: []ArgValue{
						value,
					},
				},
			},
		},
	}

	t.Run("GetFlagArgValues", func(t *testing.T) {
		v, ok := parsedData.GetFlagArgValues(flag)
		require.True(t, ok)
		require.Equal(t, value, v[0])
	})

	t.Run("GetFlagArgValue", func(t *testing.T) {
		v, ok := parsedData.GetFlagArgValue(flag)
		require.True(t, ok)
		require.Equal(t, value, v)
	})
}
