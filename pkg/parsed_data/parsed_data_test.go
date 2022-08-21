package parsed_data

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"testing"
)

func TestParsedDataGetters(t *testing.T) {
	t.Parallel()

	var pointer *ParsedData

	{
		require.Equal(t, dollyconf.CommandIDUndefined, pointer.GetCommandID())
		require.Equal(t, dollyconf.CommandUndefined, pointer.GetCommand())
		require.Nil(t, pointer.GetAgrData())
		require.Nil(t, pointer.GetFlagDataMap())
	}
	{
		pointer = NewParsedData(
			dollyconf.CommandID(gofakeit.Uint32()),
			dollyconf.Command(gofakeit.Name()),
			NewParsedArgData(nil),
			map[dollyconf.Flag]*ParsedFlagData{
				dollyconf.Flag(gofakeit.Name()): NewParsedFlagData(
					dollyconf.Flag(gofakeit.Name()),
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
			dollyconf.CommandID(gofakeit.Uint32()),
			dollyconf.Command(gofakeit.Name()),
			NewParsedArgData(nil),
			nil,
		)

		require.Nil(t, pointer.GetFlagDataMap())
	}
}

func TestGetFlagArgValuesErrors(t *testing.T) {
	t.Parallel()

	flag := dollyconf.Flag(gofakeit.Color())

	testData := []struct {
		caseName string

		parsedData *ParsedData
		flag       dollyconf.Flag

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
				FlagDataMap: map[dollyconf.Flag]*ParsedFlagData{
					flag: {
						Flag:    flag,
						ArgData: &ParsedArgData{},
					},
				},
			},
			expectedSuccess: true,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName+"_GetFlagArgValues", func(t *testing.T) {
			v, ok := td.parsedData.GetFlagArgValues(td.flag)
			require.Equal(t, 0, len(v))
			require.Equal(t, td.expectedSuccess, ok)
		})
	}
}

func TestGetFlagArgValueErrors(t *testing.T) {
	t.Parallel()

	flag := dollyconf.Flag(gofakeit.Color())

	testData := []struct {
		caseName string

		parsedData *ParsedData
		flag       dollyconf.Flag

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
				FlagDataMap: map[dollyconf.Flag]*ParsedFlagData{
					flag: {
						Flag:    flag,
						ArgData: &ParsedArgData{},
					},
				},
			},
		},
	}

	for _, td := range testData {
		t.Run(td.caseName+"_GetFlagArgValues", func(t *testing.T) {
			v, ok := td.parsedData.GetFlagArgValue(td.flag)
			require.Equal(t, 0, len(v))
			require.Equal(t, td.expectedSuccess, ok)
		})
	}
}

func TestGetFlagArgValue(t *testing.T) {
	t.Parallel()

	flag := dollyconf.Flag(gofakeit.Color())
	value := ArgValue(gofakeit.Color())

	parsedData := &ParsedData{
		FlagDataMap: map[dollyconf.Flag]*ParsedFlagData{
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
