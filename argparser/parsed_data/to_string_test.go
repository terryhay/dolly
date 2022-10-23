package parsed_data

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestArgValueToString(t *testing.T) {
	t.Parallel()

	testData := []struct {
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

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			res := td.argValue.ToString()
			require.Equal(t, string(td.argValue), res)
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
