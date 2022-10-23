package row_len_limiter

import (
	"fmt"
	"github.com/terryhay/dolly/argparser/helpdisplay/size"
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
		min:     min.ToInt(),
		optimum: optimum.ToInt(),
		max:     max.ToInt(),
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
	return size.Width(rl.min)
}

// Optimum returns optimum field
func (rl *RowLenLimit) Optimum() size.Width {
	if rl == nil {
		return 0
	}
	return size.Width(rl.optimum)
}

// Max returns max field
func (rl *RowLenLimit) Max() size.Width {
	if rl == nil {
		return 0
	}
	return size.Width(rl.max)
}

// ApplyTabShift recalculates limits for new terminal change delta
func (rl *RowLenLimit) ApplyTabShift(shift size.Width) (res RowLenLimit) {
	if rl == nil {
		return RowLenLimit{}
	}

	res.min = rl.min - shift.ToInt()
	res.optimum = rl.optimum - shift.ToInt()
	res.max = rl.max - shift.ToInt()

	if res.min < 0 {
		res.min = rl.min
		res.optimum = rl.optimum
		res.max = rl.max
	}

	return res
}

func (rl *RowLenLimit) IsValid() bool {
	return rl.Min() != 0 && rl.Optimum() != 0 && rl.Max() != 0
}

func (rl *RowLenLimit) Clone() RowLenLimit {
	return RowLenLimit{
		min:     int(rl.Min()),
		optimum: int(rl.Optimum()),
		max:     int(rl.Max()),
	}
}

func (rl *RowLenLimit) String() string {
	if rl == nil {
		return ""
	}
	return fmt.Sprintf("[min: %d; optimum: %d; max: %d]", rl.min, rl.optimum, rl.max)
}
