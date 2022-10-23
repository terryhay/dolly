package arg_parser_config

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArgumentsDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *ArgumentsDescription

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, ArgAmountTypeNoArgs, pointer.GetAmountType())
		require.Equal(t, "", pointer.GetSynopsisHelpDescription())
		require.Nil(t, pointer.GetDefaultValues())
		require.Nil(t, pointer.GetAllowedValues())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &ArgumentsDescription{
			AmountType:              ArgAmountTypeSingle,
			SynopsisHelpDescription: gofakeit.Name(),
			DefaultValues:           []string{gofakeit.Name()},
			AllowedValues: map[string]bool{
				gofakeit.Name(): true,
			},
		}

		require.Equal(t, pointer.AmountType, pointer.GetAmountType())
		require.Equal(t, pointer.SynopsisHelpDescription, pointer.GetSynopsisHelpDescription())
		require.Equal(t, pointer.DefaultValues, pointer.GetDefaultValues())
		require.Equal(t, pointer.AllowedValues, pointer.GetAllowedValues())
	})
}
