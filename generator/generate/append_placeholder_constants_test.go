package generate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	ce "github.com/terryhay/dolly/generator/config_entity"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestAppendPlaceholderConstants(t *testing.T) {
	t.Parallel()

	genDataPlacehodlersSorted := []*ce.GenComponents{
		ce.NewGenComponents(ce.GenComponentsOpt{
			PrefixID: ce.PrefixPlaceholderID,
			Name:     coty.NamePlaceholder("placehoder1"),
			Comment:  coty.InfoChapterDESCRIPTION("do something shit"),
		}),
		ce.NewGenComponents(ce.GenComponentsOpt{
			PrefixID: ce.PrefixPlaceholderID,
			Name:     coty.NamePlaceholder("placehoder2"),
			Comment:  coty.InfoChapterDESCRIPTION("do something default shit"),
		}),
	}

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, appendPlaceholderConstants(nil, genDataPlacehodlersSorted))
		require.Equal(t, 0, len(appendPlaceholderConstants(&strings.Builder{}, nil).String()))
	})

	t.Run("common", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, `
const (
    // IDPlaceholderPlacehoder1 - do something shit
    IDPlaceholderPlacehoder1 coty.IDPlaceholder = iota + 1

    // IDPlaceholderPlacehoder2 - do something default shit
    IDPlaceholderPlacehoder2
)
`, appendPlaceholderConstants(&strings.Builder{}, genDataPlacehodlersSorted).String())
	})
}
