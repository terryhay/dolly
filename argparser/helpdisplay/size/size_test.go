package size

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConvert(t *testing.T) {
	t.Parallel()

	randInt := int(gofakeit.Uint32())

	require.Equal(t, randInt, Width(randInt).ToInt())
	require.Equal(t, randInt, Height(randInt).ToInt())

	require.Equal(t, Height(0), Height(0).AddInt(-1))
	require.Equal(t, Height(1), Height(0).AddInt(1))
}
