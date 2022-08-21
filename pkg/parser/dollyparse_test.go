package parser

import (
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"testing"
)

func TestArgParser(t *testing.T) {
	res, err := Parse(
		dollyconf.NewArgParserConfig(
			dollyconf.ApplicationDescription{},
			nil,
			nil,
			nil,
			nil),
		nil)

	require.Nil(t, res)
	require.NotNil(t, err)
}
