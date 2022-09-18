package file_decorator

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileDecorator(t *testing.T) {
	t.Parallel()

	var impl *fileDecoratorImpl

	require.NotNil(t, impl.Close())
	require.NotNil(t, impl.WriteString(gofakeit.Name()))

	fileDecorator := NewFileDecorator(nil)
	require.NotNil(t, fileDecorator.Close())
	require.NotNil(t, fileDecorator.WriteString(gofakeit.Name()))
}
