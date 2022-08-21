package os_decorator_mock

import (
	osd "github.com/terryhay/dolly/internal/os_decorator"
	"os"
)

// OSDecoratorMockInit - init struct
type OSDecoratorMockInit struct {
	Args           []string
	CreateFunc     func(path string) (osd.FileDecorator, error)
	ExitFunc       func(err error, code uint)
	IsNotExistFunc func(err error) bool
	MkdirAllFunc   func(path string, perm os.FileMode) error
	StatFunc       func(path string) (os.FileInfo, error)
}

// NewOSDecoratorMock - mocked os decorator instance constructor
func NewOSDecoratorMock(init OSDecoratorMockInit) osd.OSDecorator {
	return &osDecoratorMockImpl{
		mockArgs:           init.Args,
		mockCreateFunc:     init.CreateFunc,
		mockExit:           init.ExitFunc,
		mockIsNotExistFunc: init.IsNotExistFunc,
		mockMkdirAll:       init.MkdirAllFunc,
		mockStatFunc:       init.StatFunc,
	}
}

type osDecoratorMockImpl struct {
	mockArgs           []string
	mockCreateFunc     func(path string) (osd.FileDecorator, error)
	mockExit           func(err error, code uint)
	mockIsNotExistFunc func(err error) bool
	mockMkdirAll       func(path string, perm os.FileMode) error
	mockStatFunc       func(name string) (os.FileInfo, error)
}

// GetArgs Args - returns command line arguments without application name
func (i *osDecoratorMockImpl) GetArgs() []string {
	return i.mockArgs
}

// Create - creates or truncates the named file
func (i *osDecoratorMockImpl) Create(path string) (osd.FileDecorator, error) {
	return i.mockCreateFunc(path)
}

// Exit - causes the current program to exit with the given error
func (i *osDecoratorMockImpl) Exit(err error, code uint) {
	i.mockExit(err, code)
}

// IsNotExist - checks if error is "not exist"
func (i *osDecoratorMockImpl) IsNotExist(err error) bool {
	return i.mockIsNotExistFunc(err)
}

// MkdirAll - creates a directory named path
func (i *osDecoratorMockImpl) MkdirAll(path string, perm os.FileMode) error {
	return i.mockMkdirAll(path, perm)
}

// Stat - returns a FileInfo describing the named file
func (i *osDecoratorMockImpl) Stat(name string) (os.FileInfo, error) {
	return i.mockStatFunc(name)
}
