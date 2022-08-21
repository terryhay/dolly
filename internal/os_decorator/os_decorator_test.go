package os_decorator

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"os"
	"testing"
)

func TestOSDecorator(t *testing.T) {
	t.Parallel()

	osDecorator := NewOSDecorator()

	require.True(t, len(osDecorator.GetArgs()) > 0)
	{
		fileDecorator, err := osDecorator.Create("")
		require.NotNil(t, fileDecorator)
		require.NotNil(t, err)
	}

	require.False(t, osDecorator.IsNotExist(nil))

	require.Nil(t, osDecorator.MkdirAll("\\", 7777))

	{
		fileInfo, err := osDecorator.Stat(gofakeit.Name())
		_ = fileInfo
		require.NotNil(t, err)
	}
}

func TestOSDecoratorExit(t *testing.T) {
	t.Parallel()

	fakeExit := func(int) {
		panic("os.Exit called")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	assert.PanicsWithValue(
		t,
		"os.Exit called",
		func() {
			NewOSDecorator().Exit(
				dollyerr.NewError(dollyerr.CodeUndefinedError, fmt.Errorf(gofakeit.Name())),
				dollyerr.CodeUndefinedError.ToUint())
		},
		"os.Exit was not called")
}

func TestFileDecorator(t *testing.T) {
	t.Parallel()

	var impl *fileDecoratorImpl

	require.NotNil(t, impl.Close())
	require.NotNil(t, impl.WriteString(gofakeit.Name()))

	fileDecorator := NewFileDecorator(nil)
	require.NotNil(t, fileDecorator.Close())
	require.NotNil(t, fileDecorator.WriteString(gofakeit.Name()))
}
