package paragraph_model

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBreak(t *testing.T) {
	t.Parallel()

	var b *splitter
	require.Equal(t, 0, b.indexBegin().ToInt())
	require.Equal(t, 0, b.indexEnd().ToInt())

	b = newSplitter(1, 2)
	require.Equal(t, 1, b.indexBegin().ToInt())
	require.Equal(t, 2, b.indexEnd().ToInt())
}
