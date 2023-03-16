package common_types

import (
	"fmt"
	"sort"
)

type constraintsToSlice interface {
	NameCommand | NamePlaceholder | NameFlag | InfoChapterDESCRIPTION | ArgValue
}

// ToSliceTypesSorted converts string slice to T slice
func ToSliceTypesSorted[T constraintsToSlice](from []string) []T {
	if len(from) == 0 {
		return nil
	}

	to := make([]T, 0, len(from))
	for _, v := range from {
		to = append(to, T(v))
	}

	sort.Slice(to, func(l, r int) bool {
		return to[l] < to[r]
	})

	return to
}

// ToSliceStrings converts []T to []string
func ToSliceStrings[T fmt.Stringer](from []T) []string {
	if len(from) == 0 {
		return nil
	}

	to := make([]string, 0, len(from))
	for _, s := range from {
		to = append(to, s.String())
	}

	return to
}
