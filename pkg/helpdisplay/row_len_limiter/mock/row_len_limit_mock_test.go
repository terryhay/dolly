package rowLenLimitMock

import (
	"github.com/stretchr/testify/assert"
	"github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"testing"
)

func TestRowLenLimitMock(t *testing.T) {
	t.Parallel()

	assert.Equal(t, row_len_limiter.MakeRowLenLimit(10, 14, 20), GetRowLenLimitMin())
	assert.Equal(t, row_len_limiter.MakeRowLenLimit(13, 17, 25), GetRowLenLimitForTerminalWidth25())
	assert.Equal(t, row_len_limiter.MakeDefaultRowLenLimit(), GetRowLenLimitMax())
}
