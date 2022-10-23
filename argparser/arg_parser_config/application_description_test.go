package arg_parser_config

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApplicationDescriptionGetters(t *testing.T) {
	t.Parallel()

	obj := ApplicationDescription{
		AppName:             gofakeit.Name(),
		NameHelpInfo:        gofakeit.Name(),
		DescriptionHelpInfo: []string{gofakeit.Name()},
	}

	require.Equal(t, obj.AppName, obj.GetAppName())
	require.Equal(t, obj.NameHelpInfo, obj.GetNameHelpInfo())
	require.Equal(t, obj.DescriptionHelpInfo, obj.GetDescriptionHelpInfo())
}
