package file_writer

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	fld "github.com/terryhay/dolly/generator/file_decorator"
	osd "github.com/terryhay/dolly/generator/os_decorator"
	"github.com/terryhay/dolly/utils/dollyerr"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	t.Parallel()

	mockFuncCreateErrRes := fmt.Errorf(gofakeit.Name())
	mockMkdirAllErrRes := fmt.Errorf(gofakeit.Name())

	dirPath := gofakeit.Color()

	testCases := []struct {
		caseName string

		osDecor  osd.OSDecorator
		dirPath  string
		fileBody string

		expectedErrCode dollyerr.Code
	}{
		{
			caseName: "check_dir_path_error",

			osDecor: osd.NewOSDecorator(
				&osd.Mock{
					FuncIsExist: func(path string) bool {
						return false
					},
				}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorInvalidPath,
		},
		{
			caseName: "check_dir_path_error",

			osDecor: osd.NewOSDecorator(
				&osd.Mock{
					FuncIsExist: func(path string) bool {
						return path == dirPath
					},
					FuncMkdirAll: func(path string, perm os.FileMode) error {
						return mockMkdirAllErrRes
					},
				}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorCreateDirError,
		},
		{
			caseName: "create_file_error_with_successful_create_dir",

			osDecor: osd.NewOSDecorator(
				&osd.Mock{
					FuncCreate: func(string) (fld.FileDecorator, error) {
						return nil, mockFuncCreateErrRes
					},
					FuncIsExist: func(path string) bool {
						return len(path) > 0
					},
					FuncMkdirAll: func(path string, perm os.FileMode) error {
						return nil
					},
				}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorCreateFileError,
		},
		{
			caseName: "file_create_error",

			osDecor: osd.NewOSDecorator(
				&osd.Mock{
					FuncCreate: func(path string) (fld.FileDecorator, error) {
						return nil, mockFuncCreateErrRes
					},
					FuncIsExist: func(path string) bool {
						return path == dirPath
					},
					FuncMkdirAll: func(path string, perm os.FileMode) error {
						return nil
					},
				}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorCreateFileError,
		},
		{
			caseName: "file_create_error",

			osDecor: osd.NewOSDecorator(
				&osd.Mock{
					FuncCreate: func(path string) (fld.FileDecorator, error) {
						return fld.NewFileDecorator(
								nil,
								&fld.Mock{
									FuncClose: func() error {
										return nil
									},
									FuncWriteString: func(s string) (int, error) {
										return 0, fmt.Errorf(gofakeit.Color())
									},
								}),
							nil
					},
					FuncIsExist: func(_ string) bool {
						return true
					},
					FuncMkdirAll: func(path string, perm os.FileMode) error {
						return nil
					},
				}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeFileDecoratorWriteStringError,
		},
		{
			caseName: "write_string_error",

			osDecor: osd.NewOSDecorator(
				&osd.Mock{
					FuncCreate: func(path string) (fld.FileDecorator, error) {
						return fld.NewFileDecorator(
								nil,
								&fld.Mock{
									FuncClose: func() error {
										return nil
									},
									FuncWriteString: func(s string) (int, error) {
										return 0, fmt.Errorf(gofakeit.Color())
									},
								}),
							nil
					},
					FuncIsExist: func(path string) bool {
						return true
					},
					FuncMkdirAll: func(path string, perm os.FileMode) error {
						return nil
					},
				}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeFileDecoratorWriteStringError,
		},
		{
			caseName: "file_close_error",

			osDecor: osd.NewOSDecorator(
				&osd.Mock{
					FuncCreate: func(path string) (fld.FileDecorator, error) {
						return fld.NewFileDecorator(
								nil,
								&fld.Mock{
									FuncClose: func() error {
										return fmt.Errorf(gofakeit.Color())
									},
									FuncWriteString: func(s string) (int, error) {
										return 0, nil
									},
								}),
							nil
					},
					FuncIsExist: func(path string) bool {
						return true
					},
					FuncMkdirAll: func(path string, perm os.FileMode) error {
						return nil
					},
				}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeFileDecoratorCloseError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			err := WriteFile(tc.osDecor, tc.dirPath, tc.fileBody)
			if tc.expectedErrCode == dollyerr.CodeNone {
				require.Nil(t, err)
				return
			}
			require.Equal(t, tc.expectedErrCode, err.Code())
		})
	}
}
