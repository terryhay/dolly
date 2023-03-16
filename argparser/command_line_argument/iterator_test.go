package command_line_argument

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIterator(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer Iterator

		require.True(t, pointer.IsEnded())
		require.Equal(t, ArgumentEmpty, pointer.GetArg())
		require.Equal(t, ArgumentEmpty, pointer.Next().GetArg())
	})

	t.Run("initialized", func(t *testing.T) {
		t.Parallel()

		it := MakeIterator([]string{
			RandCmdArg().String(),
			RandCmdArgSecond().String(),
		})

		require.Equal(t, RandCmdArg(), it.GetArg())

		it = it.Next()
		require.Equal(t, RandCmdArgSecond(), it.GetArg())

		it = it.Next()
		require.Equal(t, ArgumentEmpty, it.GetArg())

		// Another one Next must do not panic
		it = it.Next()
		require.Equal(t, ArgumentEmpty, it.GetArg())
	})
}
