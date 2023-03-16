package os_proxy

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/terryhay/dolly/generator/proxyes/file_proxy"
)

// ExitCode - application exit code
type ExitCode int

// Int converts ExitCode to int
func (c ExitCode) Int() int {
	return int(c)
}

// Proxy - os methods proxy interface
type Proxy interface {
	// GetArgs Argument - returns command line arguments without application name
	GetArgs() ([]string, error)

	// Create creates or truncates the named file
	Create(path string) (file_proxy.Proxy, error)

	// Exit causes the current program to exit with the given error
	Exit(code ExitCode, err error)

	// IsExist checks if path is existence
	IsExist(path string) error

	// MkdirAll creates directory named path
	MkdirAll(path string, perm os.FileMode) error

	// ReadFile loads file
	ReadFile(path string) ([]byte, error)
}

// New constructs Proxy
func New() Proxy {
	return &impl{
		opt: Opt{
			SlotGetArgs: func() []string {
				return os.Args[1:]
			},
			SlotCreate: func(path string) (file_proxy.Proxy, error) {
				file, err := os.Create(path)

				return file_proxy.New(file), err
			},
			SlotExit: func(code ExitCode, err error) {
				exitCode := 0
				if err != nil {
					err = fmt.Errorf(`parser generator: %w`, err)
					color.New(color.FgRed).Println(err)
					exitCode = code.Int()
				}
				os.Exit(exitCode)
			},
			SlotIsExist: func(path string) error {
				_, err := os.Stat(path)

				return err
			},
			SlotMkdirAll: os.MkdirAll,
			SlotReadFile: os.ReadFile,
		},
	}
}
