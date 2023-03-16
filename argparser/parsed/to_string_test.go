package parsed

import (
	"math/rand"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func TestArgValueToString(t *testing.T) {
	t.Parallel()

	tests := []struct {
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

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			res := tc.argValue.String()
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
