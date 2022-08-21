package file_writer

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	osDec "github.com/terryhay/dolly/internal/os_decorator"
	osDecMock "github.com/terryhay/dolly/internal/os_decorator/os_decorator_mock"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	t.Parallel()

	mockCreateFuncErrRes := fmt.Errorf(gofakeit.Name())
	mockMkdirAllErrRes := fmt.Errorf(gofakeit.Name())
	mockStatFuncErrRes := fmt.Errorf(gofakeit.Name())

	dirPath := gofakeit.Color()

	testData := []struct {
		caseName string

		osd      osDec.OSDecorator
		dirPath  string
		fileBody string

		expectedErrCode dollyerr.Code
	}{
		{
			caseName: "check_dir_path_error",

			osd: osDecMock.NewOSDecoratorMock(osDecMock.OSDecoratorMockInit{
				IsNotExistFunc: func(err error) bool {
					return err != nil
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					return nil, mockStatFuncErrRes
				},
			}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorInvalidGeneratePath,
		},
		{
			caseName: "check_dir_path_error",

			osd: osDecMock.NewOSDecoratorMock(osDecMock.OSDecoratorMockInit{
				IsNotExistFunc: func(err error) bool {
					return err != nil
				},
				MkdirAllFunc: func(path string, perm os.FileMode) error {
					return mockMkdirAllErrRes
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					if path == dirPath {
						return nil, nil
					}
					return nil, mockStatFuncErrRes
				},
			}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorCreateDirError,
		},
		{
			caseName: "create_file_error_with_successful_create_dir",

			osd: osDecMock.NewOSDecoratorMock(osDecMock.OSDecoratorMockInit{
				CreateFunc: func(string) (osDec.FileDecorator, error) {
					return nil, mockCreateFuncErrRes
				},
				IsNotExistFunc: func(err error) bool {
					return err != nil
				},
				MkdirAllFunc: func(path string, perm os.FileMode) error {
					return nil
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					if path == dirPath {
						return nil, nil
					}
					return nil, mockStatFuncErrRes
				},
			}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorCreateFileError,
		},
		{
			caseName: "file_create_error",

			osd: osDecMock.NewOSDecoratorMock(osDecMock.OSDecoratorMockInit{
				CreateFunc: func(path string) (osDec.FileDecorator, error) {
					return nil, mockCreateFuncErrRes
				},
				IsNotExistFunc: func(err error) bool {
					return false
				},
				MkdirAllFunc: func(path string, perm os.FileMode) error {
					return nil
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					return nil, nil
				},
			}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorCreateFileError,
		},
		{
			caseName: "file_create_error",

			osd: osDecMock.NewOSDecoratorMock(osDecMock.OSDecoratorMockInit{
				CreateFunc: func(path string) (osDec.FileDecorator, error) {
					return osDecMock.NewMockFileDecorator(
							func() error {
								return nil
							},
							func(s string) error {
								return fmt.Errorf(gofakeit.Color())
							}),
						nil
				},
				IsNotExistFunc: func(err error) bool {
					return false
				},
				MkdirAllFunc: func(path string, perm os.FileMode) error {
					return nil
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					return nil, nil
				},
			}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorWriteFileError,
		},
		{
			caseName: "file_close_error",

			osd: osDecMock.NewOSDecoratorMock(osDecMock.OSDecoratorMockInit{
				CreateFunc: func(path string) (osDec.FileDecorator, error) {
					return osDecMock.NewMockFileDecorator(
							func() error {
								return fmt.Errorf(gofakeit.Color())
							},
							func(s string) error {
								return nil
							}),
						nil
				},
				IsNotExistFunc: func(err error) bool {
					return false
				},
				MkdirAllFunc: func(path string, perm os.FileMode) error {
					return nil
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					return nil, nil
				},
			}),
			dirPath:         dirPath,
			expectedErrCode: dollyerr.CodeGeneratorFileCloseError,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err := Write(td.osd, td.dirPath, td.fileBody)
			if td.expectedErrCode == dollyerr.CodeNone {
				require.Nil(t, err)
				return
			}
			require.Equal(t, td.expectedErrCode, err.Code())
		})
	}
}
