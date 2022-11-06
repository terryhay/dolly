package dollyerr

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
)

// Code is spec Code of error
type Code uint

const (
	// CodeNone - null value, no error
	CodeNone Code = iota

	// CodeUndefinedError - undefined generator error Code
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

	// CodeFileDecoratorCloseError - file.Close method returned an error
	CodeFileDecoratorCloseError

	// CodeFileDecoratorWriteStringError - file.WriteString returned an error
	CodeFileDecoratorWriteStringError

	// CodeGeneratorInvalidPath - path is not exist
	CodeGeneratorInvalidPath

	// CodeGeneratorCreateDirError - create a dir error
	CodeGeneratorCreateDirError

	// CodeGeneratorCreateFileError - create a file error
	CodeGeneratorCreateFileError

	// CodeGeneratorNoRequiredFlag - get required flag page error
	CodeGeneratorNoRequiredFlag

	// CodeGeneratorWriteFileError - write file error
	CodeGeneratorWriteFileError

	// CodeGetConfigReadFileError - can't read yaml config file
	CodeGetConfigReadFileError

	// CodeGetConfigUnmarshalError - some unmarshal yaml config file error
	CodeGetConfigUnmarshalError

	// CodeHelpDisplayTerminalWidthLimitError - invalid calculated terminal len limit
	CodeHelpDisplayTerminalWidthLimitError

	// CodeHelpDisplayReceiverIsNilPointer - try to call method with nil receiver pointer
	CodeHelpDisplayReceiverIsNilPointer

	// CodeHelpDisplayRunError - something wrong with help page displaying
	CodeHelpDisplayRunError

	// CodeArgParserArgValueIsNotAllowed - arg value is not found in allowed values list
	CodeArgParserArgValueIsNotAllowed

	// CodeArgParserDashInFrontOfArg - argument must not contain dash in front
	CodeArgParserDashInFrontOfArg

	// CodeArgParserCheckValueAllowabilityError - generator error: try to check a value allowability by nil pointer
	CodeArgParserCheckValueAllowabilityError

	// CodeArgParserDuplicateFlags - some flag is duplicating
	CodeArgParserDuplicateFlags

	// CodeArgParserFlagMustHaveArg - some flag doesn't have arg
	CodeArgParserFlagMustHaveArg

	// CodeArgParserIsNotInitialized - looks like Init method was not called or was called with nil CmdArgSpec pointer
	CodeArgParserIsNotInitialized

	// CodeArgParserNamelessCommandUndefined - arguments are not set, but no page about nameless command in config object
	CodeArgParserNamelessCommandUndefined

	// CodeArgParserCommandDoesNotContainArgs - command doesn't contain required args
	CodeArgParserCommandDoesNotContainArgs

	// CodeArgParserRequiredFlagIsNotSet - some required flag is not set
	CodeArgParserRequiredFlagIsNotSet

	// CodeArgParserUnexpectedArg - unexpected command argument is set
	CodeArgParserUnexpectedArg

	// CodeArgParserUnexpectedFlag - unexpected flag
	CodeArgParserUnexpectedFlag

	// CodeOSDecoratorCreateError - file creation error
	CodeOSDecoratorCreateError

	// CodeOSDecoratorMkdirAllError - directory creation error
	CodeOSDecoratorMkdirAllError

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

// NewError creates Error object
func NewError(code Code, err error) *Error {
	return &Error{
		code: code,
		err:  err,
	}
}

// NewErrorIfItIs creates Error object if err argument is not a nil pointer
func NewErrorIfItIs(code Code, prefix string, err error) *Error {
	if err == nil {
		return nil
	}
	return &Error{
		code: code,
		err:  fmt.Errorf("%s: %s", prefix, err),
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
func (i *Error) Error() error {
	if i == nil {
		return nil
	}
	return i.err
}

func Append(i *Error, err error) *Error {
	if i == nil {
		return nil
	}
	i.err = multierror.Append(i.err, err)
	return i
}
