package os_decorator

import (
	fld "github.com/terryhay/dolly/internal/file_decorator"
	"os"
)

type osDecoratorImpl struct {
	funcGetArgs    func() []string
	funcCreate     func(path string) (fld.FileDecorator, error)
	funcExit       func(err error, code uint)
	funcIsNotExist func(err error) bool
	funcMkdirAll   func(path string, perm os.FileMode) error
	funcStat       func(name string) (os.FileInfo, error)
}

// GetArgs Args - returns command line arguments without application name
func (osd *osDecoratorImpl) GetArgs() []string {
	return osd.funcGetArgs()
}

// Create - creates or truncates the named file
func (osd *osDecoratorImpl) Create(path string) (fld.FileDecorator, error) {
	return osd.funcCreate(path)
}

// Exit - causes the current program to exit with the given error
func (osd *osDecoratorImpl) Exit(err error, code uint) {
	osd.funcExit(err, code)
}

// IsNotExist - checks if error is "not exist"
func (osd *osDecoratorImpl) IsNotExist(err error) bool {
	return osd.funcIsNotExist(err)
}

// MkdirAll - creates a directory named path
func (osd *osDecoratorImpl) MkdirAll(path string, perm os.FileMode) error {
	return osd.funcMkdirAll(path, perm)
}

// Stat - returns a FileInfo describing the named file
func (osd *osDecoratorImpl) Stat(name string) (os.FileInfo, error) {
	return osd.funcStat(name)
}
