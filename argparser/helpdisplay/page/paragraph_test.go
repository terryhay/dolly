package page

import (
	"github.com/stretchr/testify/require"
	rll "github.com/terryhay/dolly/argparser/helpdisplay/row_len_limiter"
	"testing"
)

func TestParagraph(t *testing.T) {
	t.Parallel()

	prg := MakeParagraph(1, `[1mappname command[0m [1m-sa[0m [4marg[0m [[1m-la[0m [4mstr[0m=val1 [val2] [4m...[0m]`)
	require.Equal(t, 1*rll.TabSize, prg.GetTabInCells())
	require.Equal(t, 49, len(prg.GetCells()))
	require.Equal(t, "    \x1b[1mappname command\x1b[0m \x1b[1m-sa\x1b[0m \x1b[4marg\x1b[0m [\x1b[1m-la\x1b[0m \x1b[4mstr\x1b[0m=val1 [val2] \x1b[4m...\x1b[0m]", prg.String())

}
