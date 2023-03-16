package parsed

import (
	"fmt"
	"strconv"

	clArg "github.com/terryhay/dolly/argparser/command_line_argument"
)

// ArgValue - string value of an argument
type ArgValue string

// ArgValueDefault - default value of ArgValue type
const ArgValueDefault ArgValue = ""

// MakeArgValue constructs ArgValue from Argument
func MakeArgValue(arg clArg.Argument) ArgValue {
	return ArgValue(arg)
}

// String implements Stringer interface
func (i ArgValue) String() string {
	return string(i)
}

// Uint8 converts ArgValue object to uint8
func (i ArgValue) Uint8() (uint8, error) {
	valueInt16, err := strconv.ParseInt(string(i), 10, 16)
	if err != nil {
		return 0, err
	}
	if valueInt16 < 0 {
		return 0, fmt.Errorf("attemption to convert a negative number \"%d\" to unsigned", valueInt16)
	}
	return uint8(valueInt16), nil
}

// Int8 converts ArgValue object to int16
func (i ArgValue) Int8() (int8, error) {
	valueInt64, err := strconv.ParseInt(string(i), 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(valueInt64), nil
}

// Uint16 converts ArgValue object to uint16
func (i ArgValue) Uint16() (uint16, error) {
	valueInt32, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, err
	}
	if valueInt32 < 0 {
		return 0, fmt.Errorf("attemption to convert a negative number \"%d\" to unsigned", valueInt32)
	}
	return uint16(valueInt32), nil
}

// Int16 converts ArgValue object to int16
func (i ArgValue) Int16() (int16, error) {
	valueInt64, err := strconv.ParseInt(string(i), 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(valueInt64), nil
}

// Uint32 converts ArgValue object to uint32
func (i ArgValue) Uint32() (uint32, error) {
	valueInt64, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, err
	}
	if valueInt64 < 0 {
		return 0, fmt.Errorf("attemption to convert a negative number \"%d\" to unsigned", valueInt64)
	}
	return uint32(valueInt64), nil
}

// Int32 converts ArgValue object to int32
func (i ArgValue) Int32() (int32, error) {
	valueInt64, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(valueInt64), nil
}

// Uint64 converts ArgValue object to uint64
func (i ArgValue) Uint64() (uint64, error) {
	valueInt64, err := strconv.ParseUint(string(i), 10, 64)
	if err != nil {
		return 0, err
	}
	return valueInt64, nil
}

// Int64 converts ArgValue object to int64
func (i ArgValue) Int64() (int64, error) {
	valueInt64, err := strconv.ParseInt(string(i), 10, 64)
	if err != nil {
		return 0, err
	}
	return valueInt64, nil
}

// Float32 converts ArgValue object to float32
func (i ArgValue) Float32() (float32, error) {
	valueFloat32, err := strconv.ParseFloat(string(i), 32)
	if err != nil {
		var defaultFloat32 float32
		return defaultFloat32, err
	}
	return float32(valueFloat32), nil
}

// Float64 converts ArgValue object to float64
func (i ArgValue) Float64() (float64, error) {
	valueFloat64, err := strconv.ParseFloat(string(i), 64)
	if err != nil {
		var defaultFloat64 float64
		return defaultFloat64, err
	}
	return valueFloat64, nil
}
