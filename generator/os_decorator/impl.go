package os_decorator

import (
	fld "github.com/terryhay/dolly/generator/file_decorator"
	"github.com/terryhay/dolly/utils/dollyerr"
	"os"
)

type osDecoratorImpl struct {
	funcGetArgs  func() []string
	funcCreate   func(path string) (fld.FileDecorator, error)
	funcExit     func(err error, code uint)
	funcIsExist  func(path string) bool
	funcMkdirAll func(path string, perm os.FileMode) error
	funcStat     func(name string) (os.FileInfo, error)
}

// GetArgs Args - returns command line arguments without application name
func (osd *osDecoratorImpl) GetArgs() []string {
	return osd.funcGetArgs()
}

// Create - creates or truncates the named file
func (osd *osDecoratorImpl) Create(path string) (fld.FileDecorator, *dollyerr.Error) {
	fileDecor, err := osd.funcCreate(path)
	return fileDecor, dollyerr.NewErrorIfItIs(dollyerr.CodeOSDecoratorCreateError, "OSDecorator.Create", err)
}

// Exit - causes the current program to exit with the given error
func (osd *osDecoratorImpl) Exit(err error, code uint) {
	osd.funcExit(err, code)
}

// IsExist - checks if path is existence
func (osd *osDecoratorImpl) IsExist(path string) bool {
	return osd.funcIsExist(path)
}

// MkdirAll - creates a directory named path
func (osd *osDecoratorImpl) MkdirAll(path string, perm os.FileMode) *dollyerr.Error {
	return dollyerr.NewErrorIfItIs(dollyerr.CodeOSDecoratorMkdirAllError, "OSDecorator.MkdirAll", osd.funcMkdirAll(path, perm))
}
