package test_tools

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckSpaces(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		caseName    string
		text        string
		expectedRes bool
	}{
		{
			caseName: "too_much_new_lines_at_the_beginning",
			text:     "\n\ntext",
		},
		{
			caseName: "too_much_new_lines",
			text:     "\ntext\n\n\ntext",
		},
		{
			caseName: "too_much_new_lines_at_the_ending",
			text:     "text\n\n\n",
		},
		{
			caseName: "too_much_spaces_at_the_beginning",
			text:     "  text",
		},
		{
			caseName: "too_much_spaces",
			text:     " text  text",
		},
		{
			caseName: "too_much_spaces_at_the_ending",
			text:     "text   ",
		},
		{
			caseName:    "true_result",
			text:        "\ntext ",
			expectedRes: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			res, _ := CheckSpaces(tc.text)
			require.Equal(t, tc.expectedRes, res)
		})
	}
}
