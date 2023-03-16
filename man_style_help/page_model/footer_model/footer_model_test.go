package footer_model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFooterModel(t *testing.T) {
	t.Parallel()

	var ftm *FooterModel
	require.Equal(t, "", ftm.GetFooterRow().String())

	ftm = NewFooterModel()
	require.Equal(t, ":", ftm.GetFooterRow().String())
}
