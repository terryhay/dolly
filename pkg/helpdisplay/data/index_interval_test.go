package data

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIndexInterval(t *testing.T) {
	t.Parallel()

	randBegin, randEnd := int(gofakeit.Int32()), int(gofakeit.Int32())
	ii := MakeIndexInterval(randBegin, randEnd)

	require.Equal(t, randBegin, ii.GetBeginIndex())
	require.Equal(t, randEnd, ii.GetEndIndex())
}
