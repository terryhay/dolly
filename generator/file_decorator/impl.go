package file_decorator

import (
	"github.com/terryhay/dolly/utils/dollyerr"
	"os"
)

type fileDecoratorImpl struct {
	file *os.File

	funcClose       func() error
	funcWriteString func(s string) (n int, err error)
}

// Close - closes the file
func (i *fileDecoratorImpl) Close() *dollyerr.Error {
	err := i.funcClose()
	return dollyerr.NewErrorIfItIs(dollyerr.CodeFileDecoratorCloseError, "fileDecorator.Close", err)

}

// WriteString - writes a string into the file
func (i *fileDecoratorImpl) WriteString(s string) *dollyerr.Error {
	_, err := i.funcWriteString(s)
	return dollyerr.NewErrorIfItIs(dollyerr.CodeFileDecoratorWriteStringError, "fileDecorator.WriteString", err)
}
