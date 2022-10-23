package file_decorator

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileDecorator(t *testing.T) {
	t.Parallel()

	funcCLoseErr := fmt.Errorf(gofakeit.Name())
	funcWriteStringErr := fmt.Errorf(gofakeit.Name())

	fileDecorator := NewFileDecorator(nil, &Mock{
		FuncClose: func() error {
			return funcCLoseErr
		},
		FuncWriteString: func(_ string) (int, error) {
			return 0, funcWriteStringErr
		},
	})

	require.Equal(t, fmt.Sprintf("fileDecorator.Close: %s", funcCLoseErr.Error()), fileDecorator.Close().Error().Error())
	require.Equal(t, fmt.Errorf("fileDecorator.WriteString: %s", funcWriteStringErr.Error()), fileDecorator.WriteString(gofakeit.Name()).Error())
}
