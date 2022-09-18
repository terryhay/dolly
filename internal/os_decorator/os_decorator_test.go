package os_decorator

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	fld "github.com/terryhay/dolly/internal/file_decorator"
	"github.com/terryhay/dolly/pkg/dollyerr"
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

	require.False(t, osDecorator.IsNotExist(nil))

	require.Nil(t, osDecorator.MkdirAll(testDir, 0744))

	{
		fileInfo, err := osDecorator.Stat(gofakeit.Name())
		_ = fileInfo
		require.NotNil(t, err)
	}

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
				dollyerr.NewError(dollyerr.CodeUndefinedError, fmt.Errorf(gofakeit.Name())),
				dollyerr.CodeUndefinedError.ToUint())
		},
		"os.Exit was not called")
}

func TestMockOSDecorator(t *testing.T) {
	t.Parallel()

	mockArgs := []string{gofakeit.Color()}
	mockCreateFuncErrRes := fmt.Errorf(gofakeit.Name())
	mockIsNotExistFuncRes := gofakeit.Bool()
	mockMkdirAllErrRes := fmt.Errorf(gofakeit.Name())
	mockStatFuncErrRes := fmt.Errorf(gofakeit.Name())

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
			FuncIsNotExist: func(err error) bool {
				return mockIsNotExistFuncRes
			},
			FuncMkdirAll: func(path string, perm os.FileMode) error {
				return mockMkdirAllErrRes
			},
			FuncStat: func(name string) (os.FileInfo, error) {
				return nil, mockStatFuncErrRes
			},
		},
	)

	require.Equal(t, mockArgs, osDecMock.GetArgs())

	_, err := osDecMock.Create("")
	require.Equal(t, mockCreateFuncErrRes, err)

	osDecMock.Exit(nil, 0)

	res := osDecMock.IsNotExist(nil)
	require.Equal(t, res, mockIsNotExistFuncRes)

	err = osDecMock.MkdirAll("", 0)
	require.Equal(t, err, mockMkdirAllErrRes)

	_, err = osDecMock.Stat("")
	require.Equal(t, err, mockStatFuncErrRes)
}
