package common_types

import (
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func TestCommonTypes(t *testing.T) {
	t.Parallel()

	str := strings.ToLower(gofakeit.Name())

	require.Equal(t, str, NameApp(str).String())
	require.Equal(t, str, InfoChapterNAME(str).String())
	require.Equal(t, str, InfoChapterDESCRIPTION(str).String())
	require.Equal(t, str, NamePlaceholder(str).String())
	require.Equal(t, str, ArgValue(str).String())
	require.Equal(t, str, NameArgHelp(str).String())
}
