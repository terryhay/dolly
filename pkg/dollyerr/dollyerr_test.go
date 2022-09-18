package dollyerr

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDollyError(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer_getters", func(t *testing.T) {
		var err *Error
		require.Equal(t, uint(CodeNone), err.Code().ToUint())
		require.Equal(t, "", err.Error())

	})

	t.Run("getters", func(t *testing.T) {
		err := fmt.Errorf(gofakeit.Name())
		dollyErr := NewError(CodeUndefinedError, err)

		require.NotNil(t, dollyErr)
		require.Equal(t, uint(CodeUndefinedError), dollyErr.Code().ToUint())
		require.Equal(t, err.Error(), dollyErr.Error())
	})

	t.Run("appending", func(t *testing.T) {
		var dollyErr *Error
		var err = fmt.Errorf(gofakeit.Name())

		dollyErr = dollyErr.AppendError(err)
		require.Nil(t, dollyErr)

		dollyErr = NewError(CodeUndefinedError, err)
		require.NotNil(t, dollyErr)
		dollyErr = dollyErr.AppendError(err)
		require.NotNil(t, dollyErr)
	})
}

func TestNewErrorIfItIs(t *testing.T) {
	t.Parallel()

	prefix := gofakeit.Name()
	dollyErr := NewErrorIfItIs(CodeUndefinedError, prefix, nil)
	require.Nil(t, dollyErr)

	err := fmt.Errorf(gofakeit.Name())
	require.Equal(t, fmt.Sprintf("%s: %s", prefix, err.Error()), NewErrorIfItIs(CodeUndefinedError, prefix, err).Error())
}
