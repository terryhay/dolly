package file_proxy

import (
	"os"
)

// Proxy - os.File methods proxy interface
type Proxy interface {
	// Close - closes the file
	Close() error

	// WriteString - writes a string into the file
	WriteString(s string) error
}

// New constructs Proxy
func New(file *os.File) Proxy {
	return &impl{
		file: file,
		opt: Opt{
			SlotClose: file.Close,
			SlotWriteString: func(s string) error {
				_, err := file.WriteString(s)
				return err
			},
		},
	}
}
