package os_proxy

import (
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/generator/proxyes/file_proxy"
	coty "github.com/terryhay/dolly/tools/common_types"
)

const testDir = "\\"

func TestProxy(t *testing.T) {
	t.Parallel()

	t.Run("nil_slots", func(t *testing.T) {
		t.Parallel()

		{
			proxy, success := (New()).(*impl)
			require.True(t, success)

			proxy.opt.SlotGetArgs = nil
			_, err := proxy.GetArgs()
			require.ErrorIs(t, err, ErrGetArgsNoImplementation)
		}
		{
			proxy, success := (New()).(*impl)
			require.True(t, success)

			proxy.opt.SlotCreate = nil
			_, err := proxy.Create(gofakeit.Name())
			require.ErrorIs(t, err, ErrCreateNoImplementation)
		}
		{
			proxy, success := (New()).(*impl)
			require.True(t, success)

			proxy.opt.SlotCreate = nil
			_, err := proxy.Create(gofakeit.Name())
			require.ErrorIs(t, err, ErrCreateNoImplementation)
		}
		{
			proxy, success := (New()).(*impl)
			require.True(t, success)

			proxy.opt.SlotIsExist = nil
			require.ErrorIs(t, proxy.IsExist(gofakeit.Name()), ErrIsExistNoImplementation)
		}
		{
			proxy, success := (New()).(*impl)
			require.True(t, success)

			proxy.opt.SlotMkdirAll = nil
			require.ErrorIs(t, proxy.MkdirAll(gofakeit.Name(), 0), ErrMkdirAllNoImplementation)
		}
		{
			proxy, success := (New()).(*impl)
			require.True(t, success)

			proxy.opt.SlotReadFile = nil
			_, err := proxy.ReadFile(gofakeit.Name())
			require.ErrorIs(t, err, ErrReadFileNoImplementation)
		}
	})

	t.Run("exit_panic", func(t *testing.T) {
		t.Parallel()

		defer func() {
			r := recover()
			require.Equal(t, r, ErrExitNoImplementation)
		}()

		var i impl
		i.Exit(0, nil)
	})

	t.Run("initialized", func(t *testing.T) {
		t.Parallel()

		proxyOS := New()
		{
			args, err := proxyOS.GetArgs()
			require.NoError(t, err)
			require.True(t, len(args) > 0)
		}
		{
			fileDecorator, err := proxyOS.Create("")
			require.NotNil(t, fileDecorator)
			require.NotNil(t, err)
		}

		require.NoError(t, proxyOS.IsExist(os.TempDir()))

		require.Nil(t, proxyOS.MkdirAll(testDir, 0744))
		require.NoError(t, proxyOS.IsExist(testDir))
		require.NoError(t, os.Remove(testDir))
	})
}

func TestProxyExit(t *testing.T) {
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
			New().Exit(ExitCode(1), coty.RandError())
		},
		"os.Exit was not called")
}

func TestMockProxy(t *testing.T) {
	t.Parallel()

	errGetArgs := coty.RandError()
	errCreate := coty.RandErrorSecond()
	errMkdirAll := coty.RandErrorThird()
	errReadFile := coty.RandErrorFourth()

	mockArgs := []string{gofakeit.Color()}
	mockErrIsExist := func() error {
		if gofakeit.Bool() {
			return errGetArgs
		}
		return nil
	}()

	mockOS := Mock(Opt{
		SlotGetArgs: func() []string {
			return mockArgs
		},
		SlotCreate: func(path string) (file_proxy.Proxy, error) {
			return nil, errCreate
		},
		SlotExit: func(code ExitCode, err error) {

		},
		SlotIsExist: func(path string) error {
			return mockErrIsExist
		},
		SlotMkdirAll: func(path string, perm os.FileMode) error {
			return errMkdirAll
		},
		SlotReadFile: func(path string) ([]byte, error) {
			return nil, errReadFile
		},
	})

	args, errArgs := mockOS.GetArgs()
	require.NoError(t, errArgs)
	require.Equal(t, mockArgs, args)

	_, err := mockOS.Create("")
	require.ErrorIs(t, err, errCreate)

	mockOS.Exit(0, nil)

	err = mockOS.IsExist("")
	switch {
	case mockErrIsExist != nil:
		require.ErrorIs(t, err, mockErrIsExist)
	default:
		require.NoError(t, err)
	}

	require.ErrorIs(t, mockOS.MkdirAll("", 0), errMkdirAll)

	_, err = mockOS.ReadFile("")
	require.ErrorIs(t, err, errReadFile)
}

func TestProxyIsExist(t *testing.T) {
	t.Parallel()

	osDecorNotMocked := New()
	require.Error(t, osDecorNotMocked.IsExist("nonexistent/path"))
}
