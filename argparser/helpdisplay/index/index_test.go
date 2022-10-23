package index

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIndex(t *testing.T) {
	t.Parallel()

	i := MakeIndex(-1)
	require.Equal(t, Null, i)
	require.Equal(t, 0, i.ToInt())

	i = Append(i, 1)
	require.Equal(t, Index(1), i)
}
