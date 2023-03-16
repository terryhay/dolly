package row_len_limiter

import (
	"errors"
	"fmt"

	"github.com/terryhay/dolly/tools/size"
)

const (
	defaultRowLenMin     = 45
	defaultRowLenOptimum = 66
	defaultRowLenMax     = 90
)

// RowLenLimit contains terminal display restrictions
type RowLenLimit struct {
	min     int
	optimum int
	max     int
}

// MakeRowLenLimit constructs RowLenLimit object in a stack
func MakeRowLenLimit(min, optimum, max size.Width) RowLenLimit {
	return RowLenLimit{
		min:     min.Int(),
		optimum: optimum.Int(),
		max:     max.Int(),
	}
}

// CloneRowLenLimit constructs RowLenLimit object like other RowLenLimit object
func CloneRowLenLimit(rl RowLenLimit) RowLenLimit {
	return RowLenLimit{
		min:     int(rl.Min()),
		optimum: int(rl.Optimum()),
		max:     int(rl.Max()),
	}
}

// MakeDefaultRowLenLimit constructs RowLenLimit in a stack with default values
func MakeDefaultRowLenLimit() RowLenLimit {
	return RowLenLimit{
		min:     defaultRowLenMin,
		optimum: defaultRowLenOptimum,
		max:     defaultRowLenMax,
	}
}

// Min returns min field
func (rl *RowLenLimit) Min() size.Width {
	if rl == nil {
		return 0
	}
	return size.MakeWidth(rl.min)
}

// Optimum returns optimum field
func (rl *RowLenLimit) Optimum() size.Width {
	if rl == nil {
		return size.WidthZero
	}
	return size.Width(rl.optimum)
}

// Max returns max field
func (rl *RowLenLimit) Max() size.Width {
	if rl == nil {
		return size.WidthZero
	}
	return size.MakeWidth(rl.max)
}

// ApplyTabShift recalculates limits for new terminal change delta
func (rl *RowLenLimit) ApplyTabShift(shift size.Width) (res RowLenLimit) {
	if rl == nil {
		return RowLenLimit{}
	}

	res.min = rl.min - shift.Int()
	res.optimum = rl.optimum - shift.Int()
	res.max = rl.max - shift.Int()

	if res.min < 0 {
		res.min = rl.min
		res.optimum = rl.optimum
		res.max = rl.max
	}

	return res
}

var (
	// ErrIsValidMin - 'min' argument is zero
	ErrIsValidMin = errors.New(`RowLenLimit.IsValid: min is 0`)

	// ErrIsValidOptimum 'optimum' argument is zero
	ErrIsValidOptimum = errors.New(`RowLenLimit.IsValid: optimum is 0`)

	// ErrIsValidMax - 'max' argument is zero
	ErrIsValidMax = errors.New(`RowLenLimit.IsValid: max is 0`)

	// ErrIsValidMinMoreThanOptimum - 'min' argument must be less or equal 'optimum' argument
	ErrIsValidMinMoreThanOptimum = errors.New(`RowLenLimit.IsValid: min must be less or equal optimum`)

	// ErrIsValidOptimumMoreThanMax - 'optimum' argument must be less or equal 'max' argument
	ErrIsValidOptimumMoreThanMax = errors.New(`RowLenLimit.IsValid: optimum must be less or equal max`)
)

// IsValid checks and return if RowLenLimit is valid
func (rl *RowLenLimit) IsValid() error {
	switch {
	case rl == nil || rl.min == 0:
		return ErrIsValidMin

	case rl.optimum == 0:
		return ErrIsValidOptimum

	case rl.max == 0:
		return ErrIsValidMax

	case rl.optimum < rl.min:
		return fmt.Errorf("%w: min is '%d', optimum is '%d", ErrIsValidMinMoreThanOptimum, rl.min, rl.optimum)

	case rl.max < rl.optimum:
		return fmt.Errorf("%w: optimum is '%d', max is '%d", ErrIsValidOptimumMoreThanMax, rl.optimum, rl.max)

	default:
		return nil
	}
}

func (rl *RowLenLimit) String() string {
	if rl == nil {
		return ""
	}
	return fmt.Sprintf("[min: %d; optimum: %d; max: %d]", rl.min, rl.optimum, rl.max)
}
