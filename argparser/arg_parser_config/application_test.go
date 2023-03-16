package arg_parser_config

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/tools/common_types"
)

func TestApplicationDescription(t *testing.T) {
	t.Parallel()

	opt := ApplicationOpt{
		AppName:         common_types.RandNameApp(),
		InfoChapterNAME: common_types.RandInfoChapterName(),
	}

	obj := MakeApplication(opt)

	require.Equal(t, opt.AppName, obj.GetAppName())
	require.Equal(t, opt.InfoChapterNAME, obj.GetNameHelpInfo())
}
