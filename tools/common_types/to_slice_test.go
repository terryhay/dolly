package common_types

import (
	"sort"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func TestToSliceTypesSorted(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, ToSliceTypesSorted[NameCommand](nil))
	})

	t.Run("common", func(t *testing.T) {
		t.Parallel()

		commands := func() []string {
			const size = 2

			uniq := make(map[string]struct{}, size)
			for len(uniq) < size {
				uniq[strings.ToLower(gofakeit.Color())] = struct{}{}
			}

			res := make([]string, 0, size)
			for v := range uniq {
				res = append(res, v)
			}

			sort.Slice(res, func(l, r int) bool {
				return res[l] > res[r]
			})

			return res
		}()

		require.Equal(t,
			[]NameCommand{
				NameCommand(commands[1]),
				NameCommand(commands[0]),
			},
			ToSliceTypesSorted[NameCommand](commands),
		)
	})
}

func TestToSliceStrings(t *testing.T) {
	t.Parallel()

	require.Nil(t, ToSliceStrings[NamePlaceholder](nil))
	require.Equal(t, []string{"command"}, ToSliceStrings([]NamePlaceholder{"command"}))
}
