package parsed_data

import (
	"fmt"
	"strconv"
)

// ArgValue - string value of an argument
type ArgValue string

// ToString - converts ArgValue object to string
func (i ArgValue) ToString() string {
	return string(i)
}

// ToUint8 - converts ArgValue object to uint8
func (i ArgValue) ToUint8() (uint8, error) {
	valueInt16, err := strconv.ParseInt(string(i), 10, 16)
	if err != nil {
		return 0, err
	}
	if valueInt16 < 0 {
		return 0, fmt.Errorf("attemption to convert a negative number \"%d\" to unsigned", valueInt16)
	}
	return uint8(valueInt16), nil
}

// ToInt8 - converts ArgValue object to int16
func (i ArgValue) ToInt8() (int8, error) {
	valueInt64, err := strconv.ParseInt(string(i), 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(valueInt64), nil
}

// ToUint16 - converts ArgValue object to uint16
func (i ArgValue) ToUint16() (uint16, error) {
	valueInt32, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, err
	}
	if valueInt32 < 0 {
		return 0, fmt.Errorf("attemption to convert a negative number \"%d\" to unsigned", valueInt32)
	}
	return uint16(valueInt32), nil
}

// ToInt16 - converts ArgValue object to int16
func (i ArgValue) ToInt16() (int16, error) {
	valueInt64, err := strconv.ParseInt(string(i), 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(valueInt64), nil
}

// ToUint32 - converts ArgValue object to uint32
func (i ArgValue) ToUint32() (uint32, error) {
	valueInt64, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, err
	}
	if valueInt64 < 0 {
		return 0, fmt.Errorf("attemption to convert a negative number \"%d\" to unsigned", valueInt64)
	}
	return uint32(valueInt64), nil
}

// ToInt32 - converts ArgValue object to int32
func (i ArgValue) ToInt32() (int32, error) {
	valueInt64, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(valueInt64), nil
}

// ToUint64 - converts ArgValue object to uint64
func (i ArgValue) ToUint64() (uint64, error) {
	valueInt64, err := strconv.ParseUint(string(i), 10, 64)
	if err != nil {
		return 0, err
	}
	return valueInt64, nil
}

// ToInt64 - converts ArgValue object to int64
func (i ArgValue) ToInt64() (int64, error) {
	valueInt64, err := strconv.ParseInt(string(i), 10, 64)
	if err != nil {
		return 0, err
	}
	return valueInt64, nil
}

// ToFloat32 - converts ArgValue object to float32
func (i ArgValue) ToFloat32() (float32, error) {
	valueFloat32, err := strconv.ParseFloat(string(i), 32)
	if err != nil {
		var defaultFloat32 float32
		return defaultFloat32, err
	}
	return float32(valueFloat32), nil
}

// ToFloat64 - converts ArgValue object to float64
func (i ArgValue) ToFloat64() (float64, error) {
	valueFloat64, err := strconv.ParseFloat(string(i), 64)
	if err != nil {
		var defaultFloat64 float64
		return defaultFloat64, err
	}
	return valueFloat64, nil
}
