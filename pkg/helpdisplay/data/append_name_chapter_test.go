package data

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"strings"
	"testing"
)

func TestCreateNameChapter(t *testing.T) {
	t.Parallel()

	randAppName := gofakeit.Name()
	randNameHelpInfo := gofakeit.Name()

	paragraphs := appendNameChapterParagraphs(make([]*Paragraph, 0), dollyconf.ApplicationDescription{
		AppName:      randAppName,
		NameHelpInfo: randNameHelpInfo,
	})

	paragraphTexts := make([]string, 0, len(paragraphs))
	for i := range paragraphs {
		paragraphTexts = append(paragraphTexts, paragraphs[i].String())
	}
	text := strings.Join(paragraphTexts, "\n")

	require.Equal(t,
		fmt.Sprintf(`[1mNAME[0m
    [1m%s[0m – %s`, randAppName, randNameHelpInfo),
		text)

}
