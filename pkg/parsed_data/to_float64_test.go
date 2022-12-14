package parsed_data

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArgValueToFloat64(t *testing.T) {
	t.Parallel()

	randNotNullValue := gofakeit.Float64()
	for randNotNullValue == 0.0 {
		randNotNullValue = gofakeit.Float64()
	}

	randPositiveValue := randNotNullValue
	if randNotNullValue < 0.0 {
		randNotNullValue *= -1.
	}

	randNegativeValue := randNotNullValue
	if randNegativeValue >= 0.0 {
		randNegativeValue *= -1.
	}

	testData := []struct {
		caseName string
		argValue ArgValue

		expectedRes float64
		expectedErr bool
	}{
		{
			caseName:    "empty_string_to_float64",
			argValue:    "",
			expectedErr: true,
		},
		{
			caseName:    "float64_null",
			argValue:    "0",
			expectedRes: 0,
		},
		{
			caseName:    "valid_positive_float64",
			argValue:    ArgValue(fmt.Sprintf("%v", randPositiveValue)),
			expectedRes: randPositiveValue,
		},
		{
			caseName:    "valid_negative_float64",
			argValue:    ArgValue(fmt.Sprintf("%v", randNegativeValue)),
			expectedRes: randNegativeValue,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			res, err := td.argValue.ToFloat64()
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
