package fmt_decorator

import "fmt"

// FmtDecorator - fmt methods decorator interface
type FmtDecorator interface {
	// Println prints line
	Println(string)
}

// New constructs FmtDecorator object
func New() FmtDecorator {
	return &impl{
		funcPrintln: func(str string) {
			fmt.Println(str)
		},
	}
}

type impl struct {
	funcPrintln func(string)
}

func (i *impl) Println(str string) {
	if i == nil {
		return
	}
	i.funcPrintln(str)
}
