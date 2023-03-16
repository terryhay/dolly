package generate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
)

func TestAppendImports(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, appendImports(nil, confYML.HelpOutToolPlainText))
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
			exp:         importsPlainText,
		},
		{
			caseName:    "man_style",
			helpOutTool: confYML.HelpOutToolManStyle,
			exp:         importsManStyle,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.exp, appendImports(&strings.Builder{}, tc.helpOutTool).String())
		})
	}
}
