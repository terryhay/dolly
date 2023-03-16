package os_proxy

import (
	"errors"
	"os"

	"github.com/terryhay/dolly/generator/proxyes/file_proxy"
)

type impl struct {
	opt Opt
}

// ErrGetArgsNoImplementation - opt.SlotGetArgs is nil pointer
var ErrGetArgsNoImplementation = errors.New("os_proxy.Proxy.SlotGetArgs: no implementation this method")

// GetArgs returns command line arguments without application name
func (i *impl) GetArgs() ([]string, error) {
	if i == nil || i.opt.SlotGetArgs == nil {
		return nil, ErrGetArgsNoImplementation
	}

	return i.opt.SlotGetArgs(), nil
}

// ErrCreateNoImplementation - opt.SlotCreate is nil pointer
var ErrCreateNoImplementation = errors.New("os_proxy.Proxy.SlotCreate: no implementation this method")

// Create creates or truncates the named file
func (i *impl) Create(path string) (file_proxy.Proxy, error) {
	if i == nil || i.opt.SlotCreate == nil {
		return nil, ErrCreateNoImplementation
	}

	return i.opt.SlotCreate(path)
}

// ErrExitNoImplementation - opt.SlotExit is nil pointer
var ErrExitNoImplementation = errors.New("os_proxy.Proxy.SlotExit: no implementation this method")

// Exit - causes the current program to exit with the given error
func (i *impl) Exit(code ExitCode, err error) {
	if i == nil || i.opt.SlotExit == nil {
		panic(ErrExitNoImplementation)
	}

	i.opt.SlotExit(code, err)
}

// ErrIsExistNoImplementation - opt.SlotIsExist is nil pointer
var ErrIsExistNoImplementation = errors.New("os_proxy.Proxy.SlotIsExist: no implementation this method")

// IsExist checks if path is existence
func (i *impl) IsExist(path string) error {
	if i == nil || i.opt.SlotIsExist == nil {
		return ErrIsExistNoImplementation
	}

	return i.opt.SlotIsExist(path)
}

// ErrMkdirAllNoImplementation - opt.SlotMkdirAll is nil pointer
var ErrMkdirAllNoImplementation = errors.New("os_proxy.Proxy.SlotMkdirAll: no implementation this method")

// MkdirAll creates a directory named path
func (i *impl) MkdirAll(path string, perm os.FileMode) error {
	if i == nil || i.opt.SlotMkdirAll == nil {
		return ErrMkdirAllNoImplementation
	}

	return i.opt.SlotMkdirAll(path, perm)
}

// ErrReadFileNoImplementation - opt.SlotReadFile is nil pointer
var ErrReadFileNoImplementation = errors.New("os_proxy.Proxy.SlotReadFile: no implementation this method")

// ReadFile loads file
func (i *impl) ReadFile(path string) ([]byte, error) {
	if i == nil || i.opt.SlotReadFile == nil {
		return nil, ErrReadFileNoImplementation
	}

	return i.opt.SlotReadFile(path)
}
