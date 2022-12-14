package parsed_data

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArgValueToInt64(t *testing.T) {
	t.Parallel()

	randNotNullValue := gofakeit.Int64()
	for randNotNullValue == 0 {
		randNotNullValue = gofakeit.Int64()
	}

	randPositiveValue := randNotNullValue
	if randPositiveValue < 0 {
		randPositiveValue *= -1
	}
	randNegativeValue := randNotNullValue
	if randNegativeValue >= 0 {
		randNegativeValue *= -1
	}

	testData := []struct {
		caseName string
		argValue ArgValue

		expectedRes int64
		expectedErr bool
	}{
		{
			caseName:    "empty_string_to_int64",
			argValue:    "",
			expectedErr: true,
		},
		{
			caseName:    "int64_null",
			argValue:    "0",
			expectedRes: 0,
		},
		{
			caseName:    "valid_positive_int64",
			argValue:    ArgValue(fmt.Sprintf("%v", randPositiveValue)),
			expectedRes: randPositiveValue,
		},
		{
			caseName:    "valid_negative_int64",
			argValue:    ArgValue(fmt.Sprintf("%v", randNegativeValue)),
			expectedRes: randNegativeValue,
		},
		{
			caseName:    "rand_positive_int64_overflow",
			argValue:    ArgValue("18446744073709551616"),
			expectedErr: true,
		},
		{
			caseName:    "rand_negative_int64_overflow",
			argValue:    ArgValue("-18446744073709551616"),
			expectedErr: true,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			res, err := td.argValue.ToInt64()
			if td.expectedErr {
				require.NotNil(t, err)
				require.Equal(t, td.expectedRes, res)

				return
			}

			require.Nil(t, err)
			require.Equal(t, td.expectedRes, res)
		})
	}
}

func TestArgValueToUint64(t *testing.T) {
	t.Parallel()

	randPositiveValue := gofakeit.Uint64()
	for randPositiveValue == 0 {
		randPositiveValue = gofakeit.Uint64()
	}

	require.True(t, randPositiveValue != 0)
	require.True(t, randPositiveValue > 0)

	testData := []struct {
		caseName string
		argValue ArgValue

		expectedRes uint64
		expectedErr bool
	}{
		{
			caseName:    "empty_string_to_int64",
			argValue:    "",
			expectedErr: true,
		},
		{
			caseName:    "int64_null",
			argValue:    "0",
			expectedRes: 0,
		},
		{
			caseName:    "valid_positive_int64",
			argValue:    ArgValue(fmt.Sprintf("%v", randPositiveValue)),
			expectedRes: randPositiveValue,
		},
		{
			caseName:    "valid_negative_int64",
			argValue:    ArgValue(fmt.Sprintf("%v", -1*int64(randPositiveValue))),
			expectedErr: true,
		},
		{
			caseName:    "rand_positive_int64_overflow",
			argValue:    ArgValue("18446744073709551616"),
			expectedErr: true,
		},
		{
			caseName:    "rand_negative_int64_overflow",
			argValue:    ArgValue("-18446744073709551616"),
			expectedErr: true,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			res, err := td.argValue.ToUint64()
			if td.expectedErr {
				require.NotNil(t, err)
				require.Equal(t, td.expectedRes, res)

				return
			}

			require.Nil(t, err)
			require.Equal(t, td.expectedRes, res)
		})
	}
}
