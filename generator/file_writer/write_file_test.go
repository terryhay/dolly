package file_writer

import (
	"os"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/generator/proxyes/file_proxy"
	"github.com/terryhay/dolly/generator/proxyes/os_proxy"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestWrite(t *testing.T) {
	t.Parallel()

	dirPath := gofakeit.Color()

	errIsExist := coty.RandError()
	errMkdirAll := coty.RandErrorSecond()
	errCreate := coty.RandErrorThird()
	errWriteString := coty.RandErrorFourth()
	errClose := coty.RandErrorFifth()

	tests := []struct {
		caseName string

		osDecor  os_proxy.Proxy
		dirPath  string
		fileBody string

		expErr error
	}{
		{
			caseName: "check_dir_path_error",

			osDecor: os_proxy.Mock(os_proxy.Opt{
				SlotIsExist: func(path string) error {
					return errIsExist
				},
			}),
			dirPath: dirPath,

			expErr: errIsExist,
		},
		{
			caseName: "check_dir_path_error",

			osDecor: os_proxy.Mock(os_proxy.Opt{
				SlotIsExist: func(path string) error {
					if path != dirPath {
						return errIsExist
					}
					return nil
				},
				SlotMkdirAll: func(path string, perm os.FileMode) error {
					return errMkdirAll
				},
			}),
			dirPath: dirPath,

			expErr: errMkdirAll,
		},
		{
			caseName: "create_file_error_with_successful_create_dir",

			osDecor: os_proxy.Mock(os_proxy.Opt{
				SlotCreate: func(string) (file_proxy.Proxy, error) {
					return nil, errCreate
				},
				SlotIsExist: func(path string) error {
					if len(path) == 0 {
						return errIsExist
					}
					return nil
				},
				SlotMkdirAll: func(path string, perm os.FileMode) error {
					return nil
				},
			}),
			dirPath: dirPath,

			expErr: errCreate,
		},
		{
			caseName: "file_create_error",

			osDecor: os_proxy.Mock(os_proxy.Opt{
				SlotCreate: func(path string) (file_proxy.Proxy, error) {
					return nil, errCreate
				},
				SlotIsExist: func(path string) error {
					if path != dirPath {
						return errIsExist
					}
					return nil
				},
				SlotMkdirAll: func(path string, perm os.FileMode) error {
					return nil
				},
			}),
			dirPath: dirPath,

			expErr: errCreate,
		},
		{
			caseName: "file_create_error",

			osDecor: os_proxy.Mock(os_proxy.Opt{
				SlotCreate: func(path string) (file_proxy.Proxy, error) {
					return file_proxy.Mock(file_proxy.Opt{
							SlotClose: func() error {
								return nil
							},
							SlotWriteString: func(s string) error {
								return errWriteString
							},
						}),
						nil
				},
				SlotIsExist: func(string) error {
					return nil
				},
				SlotMkdirAll: func(path string, perm os.FileMode) error {
					return nil
				},
			}),
			dirPath: dirPath,

			expErr: errWriteString,
		},
		{
			caseName: "write_string_error",

			osDecor: os_proxy.Mock(os_proxy.Opt{
				SlotCreate: func(path string) (file_proxy.Proxy, error) {
					return file_proxy.Mock(file_proxy.Opt{
							SlotClose: func() error {
								return nil
							},
							SlotWriteString: func(s string) error {
								return errWriteString
							},
						}),
						nil
				},
				SlotIsExist: func(path string) error {
					return nil
				},
				SlotMkdirAll: func(path string, perm os.FileMode) error {
					return nil
				},
			}),
			dirPath: dirPath,
			expErr:  errWriteString,
		},
		{
			caseName: "write_string_and_close_errors",

			osDecor: os_proxy.Mock(os_proxy.Opt{
				SlotCreate: func(path string) (file_proxy.Proxy, error) {
					return file_proxy.Mock(file_proxy.Opt{
							SlotClose: func() error {
								return errClose
							},
							SlotWriteString: func(s string) error {
								return errWriteString
							},
						}),
						nil
				},
				SlotIsExist: func(path string) error {
					return nil
				},
				SlotMkdirAll: func(path string, perm os.FileMode) error {
					return nil
				},
			}),
			dirPath: dirPath,
			expErr:  errWriteString,
		},
		{
			caseName: "file_close_error",

			osDecor: os_proxy.Mock(os_proxy.Opt{
				SlotCreate: func(path string) (file_proxy.Proxy, error) {
					return file_proxy.Mock(file_proxy.Opt{
							SlotClose: func() error {
								return errClose
							},
							SlotWriteString: func(s string) error {
								return nil
							},
						}),
						nil
				},
				SlotIsExist: func(path string) error {
					return nil
				},
				SlotMkdirAll: func(path string, perm os.FileMode) error {
					return nil
				},
			}),
			dirPath: dirPath,

			expErr: errClose,
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseName, func(t *testing.T) {
			err := WriteFile(tc.osDecor, tc.dirPath, tc.fileBody)
			if tc.expErr == nil {
				require.Nil(t, err)
				return
			}
			require.ErrorIs(t, err, tc.expErr)
		})
	}
}
