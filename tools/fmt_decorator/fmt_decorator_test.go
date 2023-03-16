package fmt_decorator

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	tools "github.com/terryhay/dolly/tools/test_tools"
)

func TestNewFmtDecorator(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *impl
		pointer.Println(gofakeit.Name())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		t.Parallel()

		fmtDec := New()
		require.NotNil(t, fmtDec)

		str := gofakeit.Name()

		require.Equal(t, str+"\n", tools.CatchStdOut(func() {
			fmtDec.Println(str)
		}))
	})

}
