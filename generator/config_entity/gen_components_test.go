package config_entity

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/tools/common_types"
)

func TestGenerateDataGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var pointer *GenComponents

		require.Equal(t, 0, len(pointer.GetNameID()))
		require.Equal(t, 0, len(pointer.GetName()))
		require.Equal(t, 0, len(pointer.GetComment()))
	})

	t.Run("valid_pointer", func(t *testing.T) {
		opt := GenComponentsOpt{
			PrefixID: PrefixNameCommand,
			Name:     common_types.RandNameCommand(),
			Comment:  common_types.RandInfoChapterDescription(),
		}

		pointer := NewGenComponents(opt)

		require.True(t, len(pointer.GetNameID()) > 0)
		require.Equal(t, opt.Name.String(), pointer.GetName().String())
		require.Equal(t, opt.Comment.String(), pointer.GetComment())
	})
}
