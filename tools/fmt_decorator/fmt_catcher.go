package fmt_decorator

// FmtCatcher contain output data from fmt decorator methods for testing
type FmtCatcher struct {
	printlnInput string
}

// NewCatcher constructs FmtCatcher object
func NewCatcher() *FmtCatcher {
	return &FmtCatcher{}
}

// Println implements FmtDecorator
func (i *FmtCatcher) Println(str string) {
	if i == nil {
		return
	}
	i.printlnInput = str
}

// GetPrintln returns data which were printed by Println
func (i *FmtCatcher) GetPrintln() string {
	if i == nil {
		return ""
	}

	return i.printlnInput
}
