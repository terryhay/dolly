package runes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunes(t *testing.T) {
	t.Parallel()

	assert.Equal(t, string(RuneBr), "\n")
	assert.Equal(t, string(RuneColon), ":")
	assert.Equal(t, string(RuneDot), ".")
	assert.Equal(t, string(RuneEsc), "\x1b")
	assert.Equal(t, string(RuneTab), "\t")
	assert.Equal(t, string(RuneLwM), "m")
	assert.Equal(t, string(RuneLwQ), "q")
	assert.Equal(t, string(RuneLwQRu), "й")
	assert.Equal(t, string(RuneUpQ), "Q")
	assert.Equal(t, string(RuneUpQRu), "Й")
	assert.Equal(t, string(RuneSpace), " ")
}
