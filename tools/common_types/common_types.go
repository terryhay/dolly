package common_types

type (
	// NameApp - string with application name
	NameApp string

	// InfoChapterNAME - string with information for NAME chapter
	InfoChapterNAME string

	// InfoChapterDESCRIPTION - string with information for DESCRIPTION chapter
	InfoChapterDESCRIPTION string

	// NamePlaceholder - flag/argument placeholder string id
	NamePlaceholder string

	// IDPlaceholder - flag/argument placeholder id
	IDPlaceholder uint8

	// ArgValue - value of argument as string
	ArgValue string
)

const (
	// AppNameUndefined - default NameApp value, don't use it
	AppNameUndefined NameApp = ""

	// InfoChapterNAMEUndefined - default InfoChapterNAME value, don't use it
	InfoChapterNAMEUndefined InfoChapterNAME = ""

	// InfoChapterDESCRIPTIONUndefined - default InfoChapterDESCRIPTION value, don't use it
	InfoChapterDESCRIPTIONUndefined InfoChapterDESCRIPTION = ""

	// NamePlaceholderUndefined - default NamePlaceholder value, don't use it
	NamePlaceholderUndefined NamePlaceholder = ""

	// ArgPlaceholderIDUndefined - default IDPlaceholder value, don't use it
	ArgPlaceholderIDUndefined IDPlaceholder = 0

	// NameArgHelpUndefined - default NameArgHelp value, don't use it
	NameArgHelpUndefined NameArgHelp = ""
)

// String implements Stringer interface
func (n NameApp) String() string {
	return string(n)
}

// String implements Stringer interface
func (i InfoChapterNAME) String() string {
	return string(i)
}

// String implements Stringer interface
func (i InfoChapterDESCRIPTION) String() string {
	return string(i)
}

// String implements Stringer interface
func (n NamePlaceholder) String() string {
	return string(n)
}

// String implements Stringer interface
func (n ArgValue) String() string {
	return string(n)
}
