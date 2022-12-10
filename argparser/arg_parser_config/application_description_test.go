package arg_parser_config

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApplicationDescriptionGetters(t *testing.T) {
	t.Parallel()

	src := ApplicationDescriptionSrc{
		AppName:             gofakeit.Name(),
		NameHelpInfo:        gofakeit.Name(),
		DescriptionHelpInfo: []string{gofakeit.Name()},
	}

	obj := src.Cast()

	require.Equal(t, src.AppName, obj.GetAppName())
	require.Equal(t, src.NameHelpInfo, obj.GetNameHelpInfo())
	require.Equal(t, src.DescriptionHelpInfo, obj.GetDescriptionHelpInfo())
}
