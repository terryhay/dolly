package file_decorator

import (
	"github.com/terryhay/dolly/utils/dollyerr"
	"os"
)

// FileDecorator - file methods decorator interface
type FileDecorator interface {
	// Close - closes the file
	Close() *dollyerr.Error

	// WriteString - writes a string into the file
	WriteString(s string) *dollyerr.Error
}

// Mock contains mock methods of initialize mocked decorator object
type Mock struct {
	FuncClose       func() error
	FuncWriteString func(s string) (n int, err error)
}

// NewFileDecorator - file decorator instance constructor
func NewFileDecorator(file *os.File, mock *Mock) FileDecorator {
	fd := &fileDecoratorImpl{
		file: file,

		funcClose:       file.Close,
		funcWriteString: file.WriteString,
	}

	if mock != nil {
		fd.funcClose = mock.FuncClose
		fd.funcWriteString = mock.FuncWriteString
	}

	return fd
}
