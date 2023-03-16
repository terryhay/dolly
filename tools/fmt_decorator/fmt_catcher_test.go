package fmt_decorator

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func TestFmtCatcher(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *FmtCatcher
		pointer.Println(gofakeit.Name())
		require.Empty(t, pointer.GetPrintln())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		t.Parallel()

		pointer := NewCatcher()

		exp := gofakeit.Name()
		pointer.Println(exp)

		require.Equal(t, exp, pointer.GetPrintln())
	})
}
