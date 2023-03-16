package index

import "github.com/brianvoe/gofakeit"

type typeRandStorage struct {
	indexes []Index
}

var randStorage = func() typeRandStorage {
	const countRandIndexes = 1

	return typeRandStorage{
		indexes: func() []Index {
			uniq := make(map[Index]struct{}, countRandIndexes)
			for len(uniq) < countRandIndexes {
				uniq[Index(gofakeit.Uint16())] = struct{}{}
			}

			indexes := make([]Index, 0, countRandIndexes)
			for i := range uniq {
				indexes = append(indexes, i)
			}

			return indexes
		}(),
	}
}()

// RandIndex returns random index (0 may be returned)
func RandIndex() Index {
	return randStorage.indexes[0]
}
