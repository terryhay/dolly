package generate

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/tools/common_types"
)

func TestAppendNameAppConstant(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		require.Nil(t, appendNameAppConstant(nil, common_types.RandNameApp()))
		require.Equal(t, 0, len(appendPlaceholderConstants(&strings.Builder{}, nil).String()))
	})

	t.Run("common", func(t *testing.T) {
		require.Equal(t, fmt.Sprintf(`
const (
    // NameApp - name of the application
    NameApp coty.NameApp = "%s"
)
`, common_types.RandNameApp()), appendNameAppConstant(&strings.Builder{}, common_types.RandNameApp()).String())
	})
}
