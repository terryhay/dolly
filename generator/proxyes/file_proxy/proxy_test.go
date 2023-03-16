package file_proxy

import (
	"testing"

	coty "github.com/terryhay/dolly/tools/common_types"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func TestProxy(t *testing.T) {
	t.Parallel()

	t.Run("nil_slots", func(t *testing.T) {
		t.Parallel()

		proxy := New(nil)
		require.Error(t, proxy.Close())
		require.Error(t, proxy.WriteString(gofakeit.Name()))

		mock := Mock(Opt{
			SlotWriteString: proxy.WriteString,
		})
		require.ErrorIs(t, mock.Close(), ErrCloseNoImplementation)

		mock = Mock(Opt{
			SlotClose: proxy.Close,
		})
		require.ErrorIs(t, mock.WriteString(gofakeit.Name()), ErrWriteStringNoImplementation)
	})

	t.Run("mock", func(t *testing.T) {
		t.Parallel()

		proxy := Mock(Opt{
			SlotClose: func() error {
				return coty.RandError()
			},
			SlotWriteString: func(string) error {
				return coty.RandErrorSecond()
			},
		})

		require.ErrorIs(t, proxy.Close(), coty.RandError())
		require.Equal(t, proxy.WriteString(gofakeit.Name()), coty.RandErrorSecond())
	})
}
