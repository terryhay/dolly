package index

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomizer(t *testing.T) {
	t.Parallel()

	require.NotPanics(t, func() { RandIndex() })
}
