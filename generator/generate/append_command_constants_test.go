package generate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	ce "github.com/terryhay/dolly/generator/config_entity"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestAppendCommandConstants(t *testing.T) {
	t.Parallel()

	genDataCommandsSorted := []*ce.GenComponents{
		ce.NewGenComponents(ce.GenComponentsOpt{
			PrefixID: ce.PrefixNameCommand,
			Name:     coty.NameCommand("command"),
			Comment:  coty.InfoChapterDESCRIPTION("do something shit"),
		}),
		ce.NewGenComponents(ce.GenComponentsOpt{
			PrefixID: ce.PrefixNameCommand,
			Name:     ce.NamelessCommandIDPostfix,
			Comment:  coty.InfoChapterDESCRIPTION("do something default shit"),
		}),
	}

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, appendCommandConstants(nil, genDataCommandsSorted))
		require.Equal(t, 0, len(appendCommandConstants(&strings.Builder{}, nil).String()))
	})

	t.Run("common", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, `
const (
    // NameCommandCommand - do something shit
    NameCommandCommand coty.NameCommand = "command"

    // NameCommandNameless - do something default shit
    NameCommandNameless coty.NameCommand = "Nameless"
)
`, appendCommandConstants(&strings.Builder{}, genDataCommandsSorted).String())
	})
}
