package file_writer

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	fld "github.com/terryhay/dolly/internal/file_decorator"
	fldMock "github.com/terryhay/dolly/internal/file_decorator/file_decorator_mock"
	osd "github.com/terryhay/dolly/internal/os_decorator"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	t.Parallel()

	mockFuncCreateErrRes := fmt.Errorf(gofakeit.Name())
	mockMkdirAllErrRes := fmt.Errorf(gofakeit.Name())
	mockFuncStatErrRes := fmt.Errorf(gofakeit.Name())

	dirPath := gofakeit.Color()

	testData := []struct {
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
					FuncIsNotExist: func(err error) bool {
						return err != nil
					},
					FuncStat: func(path string) (os.FileInfo, error) {
						return nil, mockFuncStatErrRes
					},
				}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorInvalidGeneratePath,
		},
		{
			caseName: "check_dir_path_error",

			osDecor: osd.NewOSDecorator(
				&osd.Mock{
					FuncIsNotExist: func(err error) bool {
						return err != nil
					},
					FuncMkdirAll: func(path string, perm os.FileMode) error {
						return mockMkdirAllErrRes
					},
					FuncStat: func(path string) (os.FileInfo, error) {
						if path == dirPath {
							return nil, nil
						}
						return nil, mockFuncStatErrRes
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
					FuncIsNotExist: func(err error) bool {
						return err != nil
					},
					FuncMkdirAll: func(path string, perm os.FileMode) error {
						return nil
					},
					FuncStat: func(path string) (os.FileInfo, error) {
						if path == dirPath {
							return nil, nil
						}
						return nil, mockFuncStatErrRes
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
					FuncIsNotExist: func(err error) bool {
						return false
					},
					FuncMkdirAll: func(path string, perm os.FileMode) error {
						return nil
					},
					FuncStat: func(path string) (os.FileInfo, error) {
						return nil, nil
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
						return fldMock.NewMockFileDecorator(
								func() error {
									return nil
								},
								func(s string) error {
									return fmt.Errorf(gofakeit.Color())
								}),
							nil
					},
					FuncIsNotExist: func(err error) bool {
						return false
					},
					FuncMkdirAll: func(path string, perm os.FileMode) error {
						return nil
					},
					FuncStat: func(path string) (os.FileInfo, error) {
						return nil, nil
					},
				}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorWriteFileError,
		},
		{
			caseName: "file_close_error",

			osDecor: osd.NewOSDecorator(
				&osd.Mock{
					FuncCreate: func(path string) (fld.FileDecorator, error) {
						return fldMock.NewMockFileDecorator(
								func() error {
									return fmt.Errorf(gofakeit.Color())
								},
								func(s string) error {
									return nil
								}),
							nil
					},
					FuncIsNotExist: func(err error) bool {
						return false
					},
					FuncMkdirAll: func(path string, perm os.FileMode) error {
						return nil
					},
					FuncStat: func(path string) (os.FileInfo, error) {
						return nil, nil
					},
				}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorFileCloseError,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err := Write(td.osDecor, td.dirPath, td.fileBody)
			if td.expectedErrCode == dollyerr.CodeNone {
				require.Nil(t, err)
				return
			}
			require.Equal(t, td.expectedErrCode, err.Code())
		})
	}
}
