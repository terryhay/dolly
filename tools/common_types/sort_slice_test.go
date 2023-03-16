package common_types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortSlice(t *testing.T) {
	t.Parallel()

	sliceSorted := []NameCommand{"cmd1", "cmd0", "cmd3"}
	SortSlice(sliceSorted)

	require.Equal(t, []NameCommand{"cmd0", "cmd1", "cmd3"}, sliceSorted)
}
