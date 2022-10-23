package parser

import (
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"testing"
)

func TestArgParser(t *testing.T) {
	res, err := Parse(
		apConf.NewArgParserConfig(
			apConf.ApplicationDescription{},
			nil,
			nil,
			nil,
			nil),
		nil)

	require.Nil(t, res)
	require.NotNil(t, err)
}
