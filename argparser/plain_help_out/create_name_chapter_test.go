package plain_help_out

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNameChapter(t *testing.T) {
	t.Parallel()

	randAppName := gofakeit.Name()
	randNameHelpInfo := gofakeit.Name()

	require.Equal(t,
		fmt.Sprintf(nameChapter, randAppName, randNameHelpInfo),
		createNameChapter(randAppName, randNameHelpInfo))
}
