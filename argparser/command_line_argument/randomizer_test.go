package command_line_argument

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomizer(t *testing.T) {
	t.Parallel()

	require.LessOrEqual(t, RandCmdArg(), RandCmdArgSecond())
}
