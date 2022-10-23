package id_template_data_creator

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIDTemplateDataGetters(t *testing.T) {
	t.Parallel()

	var pointer *IDTemplateData

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, "", pointer.GetID())
		require.Equal(t, "", pointer.GetNameID())
		require.Equal(t, "", pointer.GetCallName())
		require.Equal(t, "", pointer.GetComment())
	})

	t.Run("valid_pointer", func(t *testing.T) {
		id := gofakeit.Name()
		stringID := gofakeit.Name()
		callName := gofakeit.Name()
		comment := gofakeit.Name()

		pointer = NewIDTemplateData(id, stringID, callName, comment)

		require.Equal(t, id, pointer.GetID())
		require.Equal(t, stringID, pointer.GetNameID())
		require.Equal(t, callName, pointer.GetCallName())
		require.Equal(t, comment, pointer.GetComment())
	})
}
