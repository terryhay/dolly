package file_proxy

import (
	"errors"
	"os"
)

type impl struct {
	file *os.File
	opt  Opt
}

// ErrCloseNoImplementation - opt.SlotClose is nil pointer
var ErrCloseNoImplementation = errors.New("file_proxy.Proxy.Close: no implementation this method")

// Close - closes the file
func (i *impl) Close() error {
	if i == nil || i.opt.SlotClose == nil {
		return ErrCloseNoImplementation
	}

	return i.opt.SlotClose()
}

// ErrWriteStringNoImplementation - opt.SlotWriteString is nil pointer
var ErrWriteStringNoImplementation = errors.New("file_proxy.Proxy.WriteString: no implementation this method")

// WriteString - writes a string into the file
func (i *impl) WriteString(s string) error {
	if i == nil || i.opt.SlotWriteString == nil {
		return ErrWriteStringNoImplementation
	}

	return i.opt.SlotWriteString(s)
}
