package parsed

import (
	"sort"
	"strings"

	"github.com/brianvoe/gofakeit"
)

type typeRandStorage struct {
	values []ArgValue
}

var randStorage = func() typeRandStorage {
	const countValues = 3

	return typeRandStorage{
		values: func() []ArgValue {
			uniq := make(map[ArgValue]struct{}, countValues)

			for len(uniq) < countValues {
				uniq[ArgValue(strings.ToLower(gofakeit.Color()))] = struct{}{}
			}

			res := make([]ArgValue, 0, countValues)
			for v := range uniq {
				res = append(res, v)
			}

			sort.Slice(res, func(l, r int) bool {
				return res[l] < res[r]
			})

			return res
		}(),
	}
}()

// RandArgValue returns random ArgValue
func RandArgValue() ArgValue {
	return randStorage.values[0]
}

// RandArgValueSecond returns random ArgValue
func RandArgValueSecond() ArgValue {
	return randStorage.values[1]
}

// RandArgValueThird returns random ArgValue
func RandArgValueThird() ArgValue {
	return randStorage.values[2]
}
