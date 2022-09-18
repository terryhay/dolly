package file_decorator

import (
	"fmt"
	"os"
)

// FileDecorator - file methods decorator interface
type FileDecorator interface {
	// Close - closes the file
	Close() error

	// WriteString - writes a string into the file
	WriteString(s string) error
}

// NewFileDecorator - file decorator instance constructor
func NewFileDecorator(file *os.File) FileDecorator {
	return &fileDecoratorImpl{
		file: file,
	}
}

type fileDecoratorImpl struct {
	file *os.File
}

// Close - closes the file
func (i *fileDecoratorImpl) Close() error {
	if i == nil {
		return fmt.Errorf("fileDecoratorImpl: try to Close method from nil pointer")
	}
	return i.file.Close()
}

// WriteString - writes a string into the file
func (i *fileDecoratorImpl) WriteString(s string) error {
	if i == nil {
		return fmt.Errorf("fileDecoratorImpl: try to WriteString method from nil pointer")
	}
	_, err := i.file.WriteString(s)
	return err
}
