package parser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDolly(t *testing.T) {
	t.Parallel()

	t.Run("no_args", func(t *testing.T) {
		res, err := Parse(nil)

		require.Nil(t, res)
		require.NotNil(t, err)
	})

	t.Run("call_help", func(t *testing.T) {
		res, err := Parse([]string{"help"})

		require.Equal(t, CommandIDHelp, res.GetCommandID())
		require.Nil(t, err)
	})

	t.Run("success_result", func(t *testing.T) {
		res, err := Parse([]string{"-c", "config path", "-o", "out path"})

		require.NotNil(t, res)
		require.Nil(t, err)
	})
}
