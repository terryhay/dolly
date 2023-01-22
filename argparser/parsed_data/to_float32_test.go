package parsed_data

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestArgValueToFloat32(t *testing.T) {
	t.Parallel()

	randNotNullValue := gofakeit.Float32()
	for randNotNullValue == 0.0 {
		randNotNullValue = gofakeit.Float32()
	}

	randPositiveValue := randNotNullValue
	if randNotNullValue < 0.0 {
		randNotNullValue *= -1.
	}

	randNegativeValue := randNotNullValue
	if randNegativeValue >= 0.0 {
		randNegativeValue *= -1.
	}

	randFloat64PositiveValue := math.MaxFloat32 + float64(randPositiveValue)
	randFloat64NegativeValue := -1. * randFloat64PositiveValue

	testCases := []struct {
		caseName string
		argValue ArgValue

		expectedRes float32
		expectedErr bool
	}{
		{
			caseName:    "empty_string_to_float32",
			argValue:    "",
			expectedErr: true,
		},
		{
			caseName:    "float32_null",
			argValue:    "0",
			expectedRes: 0,
		},
		{
			caseName:    "valid_positive_float32",
			argValue:    ArgValue(fmt.Sprintf("%v", randPositiveValue)),
			expectedRes: randPositiveValue,
		},
		{
			caseName:    "valid_negative_float32",
			argValue:    ArgValue(fmt.Sprintf("%v", randNegativeValue)),
			expectedRes: randNegativeValue,
		},
		{
			caseName:    "rand_positive_int64_overflow",
			argValue:    ArgValue(fmt.Sprintf("%v", randFloat64PositiveValue)),
			expectedErr: true,
		},
		{
			caseName:    "rand_negative_int64_overflow",
			argValue:    ArgValue(fmt.Sprintf("%v", randFloat64NegativeValue)),
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			res, err := tc.argValue.ToFloat32()
			if tc.expectedErr {
				require.NotNil(t, err)
				require.Equal(t, tc.expectedRes, res)

				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.expectedRes, res)
		})
	}
}
