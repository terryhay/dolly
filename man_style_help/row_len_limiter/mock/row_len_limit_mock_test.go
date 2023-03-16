package rowLenLimitMock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
)

func TestRowLenLimitMock(t *testing.T) {
	t.Parallel()

	assert.Equal(t, rll.MakeRowLenLimit(10, 14, 20), GetRowLenLimitMin())
	assert.Equal(t, rll.MakeRowLenLimit(13, 17, 25), GetRowLenLimitForTerminalWidth25())
	assert.Equal(t, rll.MakeDefaultRowLenLimit(), GetRowLenLimitMax())
}
