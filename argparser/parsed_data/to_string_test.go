package parsed_data

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestArgValueToString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		caseName string
		argValue ArgValue
	}{
		{
			caseName: "empty_string",
			argValue: "",
		},
		{
			caseName: "some_string",
			argValue: ArgValue(getRandString()),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			res := tc.argValue.ToString()
			require.Equal(t, string(tc.argValue), res)
		})
	}
}

func getRandString() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, gofakeit.Uint8())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
