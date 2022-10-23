package os_decorator

import (
	"fmt"
	fld "github.com/terryhay/dolly/generator/file_decorator"
	"github.com/terryhay/dolly/utils/dollyerr"
	"os"
)

// OSDecorator - os methods decorator interface
type OSDecorator interface {
	// GetArgs Args - returns command line arguments without application name
	GetArgs() []string

	// Create - creates or truncates the named file
	Create(path string) (fld.FileDecorator, *dollyerr.Error)

	// Exit - causes the current program to exit with the given error
	Exit(err error, code uint)

	// IsExist - checks if path is existence
	IsExist(path string) bool

	// MkdirAll - creates a directory named path
	MkdirAll(path string, perm os.FileMode) *dollyerr.Error
}

// Mock contains mock methods of initialize mocked decorator object
type Mock struct {
	FuncGetArgs  func() []string
	FuncCreate   func(path string) (fld.FileDecorator, error)
	FuncExit     func(err error, code uint)
	FuncIsExist  func(path string) bool
	FuncMkdirAll func(path string, perm os.FileMode) error
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
			return fld.NewFileDecorator(file, nil), err
		},
		funcExit: func(err error, code uint) {
			exitCode := 0
			if err != nil {
				fmt.Println("parser generator: " + err.Error())
				exitCode = int(code)
			}
			os.Exit(exitCode)
		},
		funcIsExist: func(path string) bool {
			_, err := os.Stat(path)

			res := true
			if err != nil {
				res = !os.IsNotExist(err)
			}
			return res
		},
		funcMkdirAll: os.MkdirAll,
		funcStat:     os.Stat,
	}
	if mock != nil {
		osd = &osDecoratorImpl{
			funcGetArgs:  mock.FuncGetArgs,
			funcCreate:   mock.FuncCreate,
			funcExit:     mock.FuncExit,
			funcIsExist:  mock.FuncIsExist,
			funcMkdirAll: mock.FuncMkdirAll,
		}
	}
	return osd
}
