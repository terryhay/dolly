package os_decorator

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	fld "github.com/terryhay/dolly/generator/file_decorator"
	"github.com/terryhay/dolly/utils/dollyerr"
	"os"
	"testing"
)

const testDir = "\\"

func TestOSDecorator(t *testing.T) {
	t.Parallel()

	osDecorator := NewOSDecorator(nil)

	require.True(t, len(osDecorator.GetArgs()) > 0)
	{
		fileDecorator, err := osDecorator.Create("")
		require.NotNil(t, fileDecorator)
		require.NotNil(t, err)
	}

	require.True(t, osDecorator.IsExist(os.TempDir()))

	require.Nil(t, osDecorator.MkdirAll(testDir, 0744))

	require.NoError(t, os.Remove(testDir))
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
			NewOSDecorator(nil).Exit(
				dollyerr.NewError(dollyerr.CodeUndefinedError, fmt.Errorf(gofakeit.Name())).Error(),
				dollyerr.CodeUndefinedError.ToUint())
		},
		"os.Exit was not called")
}

func TestMockOSDecorator(t *testing.T) {
	t.Parallel()

	mockArgs := []string{gofakeit.Color()}
	mockCreateFuncErrRes := fmt.Errorf(gofakeit.Name())
	mockIsExistFuncRes := gofakeit.Bool()
	mockMkdirAllErrRes := fmt.Errorf(gofakeit.Name())

	osDecMock := NewOSDecorator(
		&Mock{
			FuncGetArgs: func() []string {
				return mockArgs
			},
			FuncCreate: func(path string) (fld.FileDecorator, error) {
				return nil, mockCreateFuncErrRes
			},
			FuncExit: func(err error, code uint) {

			},
			FuncIsExist: func(path string) bool {
				return mockIsExistFuncRes
			},
			FuncMkdirAll: func(path string, perm os.FileMode) error {
				return mockMkdirAllErrRes
			},
		},
	)

	require.Equal(t, mockArgs, osDecMock.GetArgs())

	_, err := osDecMock.Create("")
	require.Equal(t, fmt.Errorf("OSDecorator.Create: %s", mockCreateFuncErrRes), err.Error())

	osDecMock.Exit(nil, 0)

	res := osDecMock.IsExist("")
	require.Equal(t, res, mockIsExistFuncRes)

	err = osDecMock.MkdirAll("", 0)
	require.Equal(t, fmt.Errorf("OSDecorator.MkdirAll: %s", mockMkdirAllErrRes), err.Error())
}

func TestOSDecoratorIsExist(t *testing.T) {
	t.Parallel()

	osDecorNotMocked := NewOSDecorator(nil)
	require.False(t, osDecorNotMocked.IsExist("nonexistent/path"))
}
