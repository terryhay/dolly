package generate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
)

func TestAppendHelpOutEntity(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		require.Nil(t, appendHelpOutEntity(nil, confYML.HelpOutToolPlainText))
	})

	tests := []struct {
		caseName    string
		helpOutTool confYML.HelpOutTool

		exp string
	}{
		{
			caseName:    "undefined",
			helpOutTool: confYML.HelpOutToolUndefined,
		},
		{
			caseName:    "plain_text",
			helpOutTool: confYML.HelpOutToolPlainText,
			exp:         helpOutEntityPlainText,
		},
		{
			caseName:    "man_style",
			helpOutTool: confYML.HelpOutToolManStyle,
			exp:         helpOutEntityManStyle,
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseName, func(t *testing.T) {
			require.Equal(t, tc.exp, appendHelpOutEntity(&strings.Builder{}, tc.helpOutTool).String())
		})
	}
}
