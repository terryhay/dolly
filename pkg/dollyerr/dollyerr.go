package dollyerr

import "github.com/hashicorp/go-multierror"

// Code is spec Code of error
type Code uint

const (
	// CodeNone - null value, no error
	CodeNone Code = iota

	// CodeUndefinedError - undefined internal error Code
	CodeUndefinedError

	// CodeConfigContainsDuplicateCommands - some command is duplicating
	CodeConfigContainsDuplicateCommands

	// CodeConfigContainsDuplicateFlags - some flag is duplicating
	CodeConfigContainsDuplicateFlags

	// CodeConfigDefaultValueIsNotAllowed - some default value is not allowed
	CodeConfigDefaultValueIsNotAllowed

	// CodeConfigFlagIsNotUsedInCommands - some flag is described, but not used in commands descriptions
	CodeConfigFlagIsNotUsedInCommands

	// CodeConfigUndefinedFlag - some flag is undefined in flag description list of yaml config file
	CodeConfigUndefinedFlag

	// CodeConfigIncorrectCharacterInFlagName - flag contain incorrect character in its name
	CodeConfigIncorrectCharacterInFlagName

	// CodeConfigIncorrectFlagLen - some flag has an empty or too long call name
	CodeConfigIncorrectFlagLen

	// CodeConfigFlagMustHaveDashInFront - all flag call names must have a dash in front
	CodeConfigFlagMustHaveDashInFront

	// CodeConfigUnexpectedDefaultValue - this set amount type description "single" if you want to use default values logic
	CodeConfigUnexpectedDefaultValue

	// CodeCantFindFlagNameInGroupSpec - unexpected flag name for determine using flag group
	CodeCantFindFlagNameInGroupSpec

	// CodeGeneratorInvalidGeneratePath - path is not exist
	CodeGeneratorInvalidGeneratePath

	// CodeGeneratorFileCloseError - file close error
	CodeGeneratorFileCloseError

	// CodeGeneratorCreateDirError - create a dir error
	CodeGeneratorCreateDirError

	// CodeGeneratorCreateFileError - create a file error
	CodeGeneratorCreateFileError

	// CodeGeneratorNoRequiredFlag - get required flag data error
	CodeGeneratorNoRequiredFlag

	// CodeGeneratorWriteFileError - write file error
	CodeGeneratorWriteFileError

	// CodeGetConfigReadFileError - can't read yaml config file
	CodeGetConfigReadFileError

	// CodeGetConfigUnmarshalError - some unmarshal yaml config file error
	CodeGetConfigUnmarshalError

	// CodeHelpDisplayTerminalWidthLimitError - invalid calculated terminal len limit
	CodeHelpDisplayTerminalWidthLimitError

	// CodeHelpDisplayInvalidTerminalHeightArgument - invalid terminal height argument
	CodeHelpDisplayInvalidTerminalHeightArgument

	// CodeHelpDisplayRenderError - error during help page rendering
	CodeHelpDisplayRenderError

	// CodeHelpDisplayRunError - something wrong with help page displaying
	CodeHelpDisplayRunError

	// CodeArgParserArgValueIsNotAllowed - arg value is not found in allowed values list
	CodeArgParserArgValueIsNotAllowed

	// CodeArgParserDashInFrontOfArg - argument must not contain dash in front
	CodeArgParserDashInFrontOfArg

	// CodeArgParserCheckValueAllowabilityError - internal error: try to check a value allowability by nil pointer
	CodeArgParserCheckValueAllowabilityError

	// CodeArgParserDuplicateFlags - some flag is duplicating
	CodeArgParserDuplicateFlags

	// CodeArgParserFlagMustHaveArg - some flag doesn't have arg
	CodeArgParserFlagMustHaveArg

	// CodeArgParserIsNotInitialized - looks like Init method was not called or was called with nil CmdArgSpec pointer
	CodeArgParserIsNotInitialized

	// CodeArgParserNamelessCommandUndefined - arguments are not set, but no data about nameless command in config object
	CodeArgParserNamelessCommandUndefined

	// CodeArgParserCommandDoesNotContainArgs - command doesn't contain required args
	CodeArgParserCommandDoesNotContainArgs

	// CodeArgParserRequiredFlagIsNotSet - some required flag is not set
	CodeArgParserRequiredFlagIsNotSet

	// CodeArgParserUnexpectedArg - unexpected command argument is set
	CodeArgParserUnexpectedArg

	// CodeArgParserUnexpectedFlag - unexpected flag
	CodeArgParserUnexpectedFlag

	// CodeTermBoxDecoratorInitError - can't init a termbox decorator
	CodeTermBoxDecoratorInitError

	// CodeTermBoxDecoratorFlushError - termbox Flush() method returned an error
	CodeTermBoxDecoratorFlushError

	// CodeTermBoxDecoratorClearError - termbox Clear() method returned an error
	CodeTermBoxDecoratorClearError
)

// ToUint converts Code type to uint
func (c Code) ToUint() uint {
	return uint(c)
}

// Error is detail of parser work error
type Error struct {
	code Code
	err  error
}

// NewError creates Error object and returns pointer
func NewError(code Code, err error) *Error {
	return &Error{
		code: code,
		err:  err,
	}
}

// Code returns code of error, you must check if error == nil before
func (i *Error) Code() Code {
	if i == nil {
		return CodeNone
	}
	return i.code
}

// Error decorates standard error interface
func (i *Error) Error() string {
	if i == nil {
		return ""
	}
	return i.err.Error()
}

func (i *Error) AppendError(err error) *Error {
	if i == nil {
		return nil
	}
	i.err = multierror.Append(i.err, err)
	return i
}
