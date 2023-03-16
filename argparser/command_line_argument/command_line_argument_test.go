package command_line_argument

import (
	"testing"

	"github.com/stretchr/testify/require"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestArgumentMethods(t *testing.T) {
	t.Parallel()

	t.Run("String", func(t *testing.T) {
		t.Parallel()

		var arg Argument
		require.Empty(t, arg.String())

		arg = RandCmdArg()
		require.Equal(t, string(arg), arg.String())
	})

	t.Run("NameCommand", func(t *testing.T) {
		t.Parallel()

		var arg Argument
		require.Empty(t, arg.ToNameCommand())

		arg = RandCmdArg()
		require.Equal(t, coty.NameCommand(arg), arg.ToNameCommand())
	})

	t.Run("NameFlag", func(t *testing.T) {
		t.Parallel()

		var arg Argument
		require.Empty(t, arg.ToNameFlag())

		arg = RandCmdArg()
		require.Equal(t, coty.NameFlag(arg), arg.ToNameFlag())
	})

	t.Run("IsValid", func(t *testing.T) {
		t.Parallel()

		var arg Argument
		require.False(t, arg.IsValid())

		require.True(t, RandCmdArg().IsValid())
	})

	t.Run("IsFlag", func(t *testing.T) {
		t.Parallel()

		var arg Argument
		require.False(t, arg.IsFlag())

		require.False(t, RandCmdArg().IsFlag())
		require.True(t, Argument(coty.RandNameFlagOneLetter()).IsFlag())
		require.True(t, Argument(coty.RandNameFlagShort()).IsFlag())
		require.True(t, Argument(coty.RandNameFlagLong()).IsFlag())
	})
}
