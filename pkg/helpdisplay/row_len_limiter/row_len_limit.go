package row_len_limiter

const (
	defaultRowLenMin     = 45
	defaultRowLenOptimum = 66
	defaultRowLenMax     = 90
)

// RowSize - size of a terminal display row (not a full terminal)
type RowSize uint8

// ToInt converts RowSize to int
func (i RowSize) ToInt() int {
	return int(i)
}

// RowLenLimit contains terminal display restrictions
type RowLenLimit struct {
	min     int
	optimum int
	max     int
}

// MakeRowLenLimit constructs RowLenLimit object in a stack
func MakeRowLenLimit(min, optimum, max RowSize) RowLenLimit {
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
func (rl *RowLenLimit) Min() RowSize {
	if rl == nil {
		return 0
	}
	return RowSize(rl.min)
}

// Optimum returns optimum field
func (rl *RowLenLimit) Optimum() RowSize {
	if rl == nil {
		return 0
	}
	return RowSize(rl.optimum)
}

// Max returns max field
func (rl *RowLenLimit) Max() RowSize {
	if rl == nil {
		return 0
	}
	return RowSize(rl.max)
}

// ApplyTabShift recalculates limits for new terminal change delta
func (rl *RowLenLimit) ApplyTabShift(shift RowSize) (res RowLenLimit) {
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
