package common_types

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// SortSlice sorts slice of alias types
func SortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}
