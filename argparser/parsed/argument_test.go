package parsed

import (
	"testing"

	"github.com/stretchr/testify/require"
	clArg "github.com/terryhay/dolly/argparser/command_line_argument"
)

func TestArgumentGetters(t *testing.T) {
	t.Parallel()

	var pointer *Argument
	require.Nil(t, pointer.GetArgValues())

	opt := ArgumentOpt{
		ArgValues: []ArgValue{
			MakeArgValue(clArg.RandCmdArg()),
			MakeArgValue(clArg.RandCmdArgSecond()),
		},
	}
	pointer = MakeArgument(&opt)
	require.Equal(t, opt.ArgValues, pointer.GetArgValues())
}
