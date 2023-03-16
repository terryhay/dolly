package os_proxy

import (
	"os"

	"github.com/terryhay/dolly/generator/proxyes/file_proxy"
)

// Opt contains mock methods of initialize mocked decorator object
type Opt struct {
	SlotGetArgs  func() []string
	SlotCreate   func(path string) (file_proxy.Proxy, error)
	SlotExit     func(code ExitCode, err error)
	SlotIsExist  func(path string) error
	SlotMkdirAll func(path string, perm os.FileMode) error
	SlotReadFile func(path string) ([]byte, error)
}

// Mock constructs mocked Proxy
func Mock(opt Opt) Proxy {
	return &impl{
		opt: opt,
	}
}
