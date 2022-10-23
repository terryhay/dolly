package parsed_data

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestArgValueToInt8(t *testing.T) {
	t.Parallel()

	randNotNullValue := gofakeit.Int8()
	for randNotNullValue == 0 {
		randNotNullValue = gofakeit.Int8()
	}

	randPositiveValue := randNotNullValue
	if randPositiveValue < 0 {
		randPositiveValue *= -1
	}
	randNegativeValue := randNotNullValue
	if randNegativeValue >= 0 {
		randNegativeValue *= -1
	}

	require.True(t, randNotNullValue != 0)
	require.True(t, randPositiveValue > 0)
	require.True(t, randNegativeValue < 0)

	randInt64PositiveValue := int64(math.MaxUint32) + int64(randPositiveValue)
	randInt64NegativeValue := -1 * randInt64PositiveValue

	testData := []struct {
		caseName string
		argValue ArgValue

		expectedRes int8
		expectedErr bool
	}{
		{
			caseName:    "empty_string_to_int8",
			argValue:    "",
			expectedErr: true,
		},
		{
			caseName:    "int8_null",
			argValue:    "0",
			expectedRes: 0,
		},
		{
			caseName:    "valid_positive_int8",
			argValue:    ArgValue(fmt.Sprintf("%d", randPositiveValue)),
			expectedRes: randPositiveValue,
		},
		{
			caseName:    "valid_negative_int8",
			argValue:    ArgValue(fmt.Sprintf("%d", randNegativeValue)),
			expectedRes: randNegativeValue,
		},
		{
			caseName:    "rand_positive_int64_overflow",
			argValue:    ArgValue(fmt.Sprintf("%d", randInt64PositiveValue)),
			expectedErr: true,
		},
		{
			caseName:    "rand_negative_int64_overflow",
			argValue:    ArgValue(fmt.Sprintf("%d", randInt64NegativeValue)),
			expectedErr: true,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			res, err := td.argValue.ToInt8()
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

func TestArgValueToUint8(t *testing.T) {
	t.Parallel()

	randPositiveValue := gofakeit.Uint8()
	for randPositiveValue == 0 {
		randPositiveValue = gofakeit.Uint8()
	}

	require.True(t, randPositiveValue != 0)
	require.True(t, randPositiveValue > 0)

	randInt64PositiveValue := int64(math.MaxUint32) + int64(randPositiveValue)
	randInt64NegativeValue := -1 * randInt64PositiveValue

	testData := []struct {
		caseName string
		argValue ArgValue

		expectedRes uint8
		expectedErr bool
	}{
		{
			caseName:    "empty_string_to_int8",
			argValue:    "",
			expectedErr: true,
		},
		{
			caseName:    "int8_null",
			argValue:    "0",
			expectedRes: 0,
		},
		{
			caseName:    "valid_positive_int8",
			argValue:    ArgValue(fmt.Sprintf("%d", randPositiveValue)),
			expectedRes: randPositiveValue,
		},
		{
			caseName:    "valid_negative_int8",
			argValue:    ArgValue(fmt.Sprintf("%d", -1*int64(randPositiveValue))),
			expectedErr: true,
		},
		{
			caseName:    "rand_positive_int64_overflow",
			argValue:    ArgValue(fmt.Sprintf("%d", randInt64PositiveValue)),
			expectedErr: true,
		},
		{
			caseName:    "rand_negative_int64_overflow",
			argValue:    ArgValue(fmt.Sprintf("%d", randInt64NegativeValue)),
			expectedErr: true,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			res, err := td.argValue.ToUint8()
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