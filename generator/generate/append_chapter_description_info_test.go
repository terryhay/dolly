package generate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
)

func TestAppendChapterDescriptionInfo(t *testing.T) {
	t.Parallel()

	infoChapterDescription := []coty.InfoChapterDESCRIPTION{
		"first description string",
		"second description string",
	}

	t.Run("nil", func(t *testing.T) {
		require.Nil(t, appendChapterDescriptionInfo(nil, size.WidthZero, infoChapterDescription))
		require.Equal(t, 0, len(appendChapterDescriptionInfo(&strings.Builder{}, size.WidthZero, nil).String()))
	})

	t.Run("common", func(t *testing.T) {
		require.Equal(t, `
        HelpInfoChapterDESCRIPTION: []string{
            "first description string"
            "second description string"
        }`, appendChapterDescriptionInfo(&strings.Builder{}, size.WidthTab+size.WidthTab, infoChapterDescription).String())
	})
}
