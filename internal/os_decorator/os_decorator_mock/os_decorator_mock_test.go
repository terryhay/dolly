package os_decorator_mock

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/internal/os_decorator"
	"os"
	"testing"
)

func TestNewMockOSDecorator(t *testing.T) {
	t.Parallel()

	mockArgs := []string{gofakeit.Color()}
	mockCreateFuncErrRes := fmt.Errorf(gofakeit.Name())
	mockIsNotExistFuncRes := gofakeit.Bool()
	mockMkdirAllErrRes := fmt.Errorf(gofakeit.Name())
	mockStatFuncErrRes := fmt.Errorf(gofakeit.Name())

	osDecMock := NewOSDecoratorMock(
		OSDecoratorMockInit{
			Args: mockArgs,
			CreateFunc: func(path string) (os_decorator.FileDecorator, error) {
				return nil, mockCreateFuncErrRes
			},
			ExitFunc: func(err error, code uint) {

			},
			IsNotExistFunc: func(err error) bool {
				return mockIsNotExistFuncRes
			},
			MkdirAllFunc: func(path string, perm os.FileMode) error {
				return mockMkdirAllErrRes
			},
			StatFunc: func(name string) (os.FileInfo, error) {
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

func TestMockFileDecorator(t *testing.T) {
	t.Parallel()

	mockCloseErrRes := fmt.Errorf(gofakeit.Name())
	mockWriteStringErrRes := fmt.Errorf(gofakeit.Name())

	mockFileDecorator := NewMockFileDecorator(
		func() error {
			return mockCloseErrRes
		},
		func(s string) error {
			return mockWriteStringErrRes
		})

	err := mockFileDecorator.Close()
	require.Equal(t, mockCloseErrRes, err)

	err = mockFileDecorator.WriteString("")
	require.Equal(t, mockWriteStringErrRes, err)
}
