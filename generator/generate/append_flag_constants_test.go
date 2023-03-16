package generate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	ce "github.com/terryhay/dolly/generator/config_entity"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestAppendFlagConstants(t *testing.T) {
	t.Parallel()

	genDataFlagsSorted := []*ce.GenComponents{
		ce.NewGenComponents(ce.GenComponentsOpt{
			PrefixID: ce.PrefixFlagName,
			Name:     coty.NameFlag("-f1"),
			Comment:  coty.InfoChapterDESCRIPTION("do something shit"),
		}),
		ce.NewGenComponents(ce.GenComponentsOpt{
			PrefixID: ce.PrefixFlagName,
			Name:     coty.NameFlag("-f2"),
			Comment:  coty.InfoChapterDESCRIPTION("do something default shit"),
		}),
	}

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, appendFlagConstants(nil, genDataFlagsSorted))
		require.Equal(t, 0, len(appendFlagConstants(&strings.Builder{}, nil).String()))
	})

	t.Run("common", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, `
const (
    // NameFlagF1 - do something shit
    NameFlagF1 coty.NameFlag = "-f1"

    // NameFlagF2 - do something default shit
    NameFlagF2 coty.NameFlag = "-f2"
)
`, appendFlagConstants(&strings.Builder{}, genDataFlagsSorted).String())
	})
}
