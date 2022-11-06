package footer_model

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFooterModel(t *testing.T) {
	t.Parallel()

	var ftm *FooterModel
	require.Equal(t, "", ftm.GetFooterRow().String())

	ftm = NewFooterModel()
	require.Equal(t, ":", ftm.GetFooterRow().String())
}
