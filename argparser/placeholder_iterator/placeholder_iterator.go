package placeholder_iterator

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/tools/index"
)

// Iterator contains data for iterating by command arg places
type Iterator struct {
	placeholders []*apConf.Placeholder
	index        index.Index
	started      bool
}

// Make constructs Iterator object on stack
func Make(placeholders []*apConf.Placeholder) Iterator {
	return Iterator{
		placeholders: placeholders,
	}
}

// IsEnded returns if iterator is ended
func (pi *Iterator) IsEnded() bool {
	if pi == nil {
		return true
	}
	return pi.index.Int() >= len(pi.placeholders)
}

// Get returns Placeholder
func (pi *Iterator) Get() *apConf.Placeholder {
	if pi == nil || len(pi.placeholders) <= pi.index.Int() {
		return nil
	}
	return pi.placeholders[pi.index.Int()]
}

// Next increments an index and return next Placeholder
func (pi *Iterator) Next() *apConf.Placeholder {
	if pi == nil || len(pi.placeholders) == 0 {
		return nil
	}

	if !pi.started {
		pi.started = true
		return pi.placeholders[pi.index.Int()]
	}

	pi.index = index.Append(pi.index, 1)
	if pi.index.Int() >= len(pi.placeholders) {
		return nil
	}
	return pi.placeholders[pi.index.Int()]
}
