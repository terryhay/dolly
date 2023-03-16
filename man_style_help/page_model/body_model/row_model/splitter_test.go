package row_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBreak(t *testing.T) {
	t.Parallel()

	var b *splitter
	require.Equal(t, 0, b.indexBegin().Int())
	require.Equal(t, 0, b.indexEnd().Int())

	b = newSplitter(1, 2)
	require.Equal(t, 1, b.indexBegin().Int())
	require.Equal(t, 2, b.indexEnd().Int())
}
