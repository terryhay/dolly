package os_decorator

import (
	"fmt"
	fld "github.com/terryhay/dolly/internal/file_decorator"
	"os"
)

// OSDecorator - os methods decorator interface
type OSDecorator interface {
	// GetArgs Args - returns command line arguments without application name
	GetArgs() []string

	// Create - creates or truncates the named file
	Create(path string) (fld.FileDecorator, error)

	// Exit - causes the current program to exit with the given error
	Exit(err error, code uint)

	// IsNotExist - checks if error is "not exist"
	IsNotExist(err error) bool

	// MkdirAll - creates a directory named path
	MkdirAll(path string, perm os.FileMode) error

	// Stat - returns a FileInfo describing the named file
	Stat(name string) (os.FileInfo, error)
}

// Mock contains mock methods of initialize mocked decorator object
type Mock struct {
	FuncGetArgs    func() []string
	FuncCreate     func(path string) (fld.FileDecorator, error)
	FuncExit       func(err error, code uint)
	FuncIsNotExist func(err error) bool
	FuncMkdirAll   func(path string, perm os.FileMode) error
	FuncStat       func(name string) (os.FileInfo, error)
}

// NewOSDecorator constructs a new os decorator object.
// You can mock it by mean not nil mock argument
func NewOSDecorator(mock *Mock) OSDecorator {
	osd := &osDecoratorImpl{
		funcGetArgs: func() []string {
			return os.Args[1:]
		},
		funcCreate: func(path string) (fld.FileDecorator, error) {
			file, err := os.Create(path)
			return fld.NewFileDecorator(file), err
		},
		funcExit: func(err error, code uint) {
			exitCode := 0
			if err != nil {
				fmt.Println("parser generator: " + err.Error())
				exitCode = int(code)
			}
			os.Exit(exitCode)
		},
		funcIsNotExist: os.IsNotExist,
		funcMkdirAll:   os.MkdirAll,
		funcStat:       os.Stat,
	}
	if mock != nil {
		osd = &osDecoratorImpl{
			funcGetArgs:    mock.FuncGetArgs,
			funcCreate:     mock.FuncCreate,
			funcExit:       mock.FuncExit,
			funcIsNotExist: mock.FuncIsNotExist,
			funcMkdirAll:   mock.FuncMkdirAll,
			funcStat:       mock.FuncStat,
		}
	}
	return osd
}
