package parsed_data

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParsedArgDataGetters(t *testing.T) {
	t.Parallel()

	var pointer *ParsedArgData
	require.Nil(t, pointer.GetArgValues())

	pointer = NewParsedArgData([]ArgValue{ArgValue(gofakeit.Name())})
	require.Equal(t, pointer.ArgValues, pointer.GetArgValues())
}
