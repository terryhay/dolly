package parsed

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomization(t *testing.T) {
	t.Parallel()

	require.Less(t, RandArgValue(), RandArgValueSecond())
}
