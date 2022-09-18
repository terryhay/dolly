package models

import (
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/helpdisplay/row"
	"testing"
)

func TestRowIteratorErrors(t *testing.T) {
	t.Parallel()

	t.Run("check_nil_pointer", func(t *testing.T) {
		var it *RowIterator

		require.True(t, it.End())
		require.Equal(t, row.Row{}, it.Row())
		require.Error(t, it.Next())
	})

	t.Run("call_Next_from_end_iterator", func(t *testing.T) {
		var it RowIterator
		require.Error(t, it.Next())
	})

	t.Run("", func(t *testing.T) {
		it := RowIterator{
			model:          &PageModel{},
			ReverseCounter: 1,
		}
		require.Error(t, it.Next())
	})
}
