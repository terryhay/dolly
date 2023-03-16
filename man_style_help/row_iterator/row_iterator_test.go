package row_iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	pgm "github.com/terryhay/dolly/man_style_help/page_model"
	"github.com/terryhay/dolly/man_style_help/row"
	rllMock "github.com/terryhay/dolly/man_style_help/row_len_limiter/mock"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
)

func TestRowIterator(t *testing.T) {
	t.Parallel()

	t.Run("ending_error", func(t *testing.T) {
		var ri RowIterator
		ri.Next()
		require.True(t, ri.End())
	})

	t.Run("get_paragraph_error", func(t *testing.T) {
		ri := RowIterator{
			model:          &pgm.PageModel{},
			reverseCounter: 1,
		}
		ri.Next()
	})

	t.Run("mocking", func(t *testing.T) {
		m := MockRowIterator(Mock{
			Model: &pgm.PageModel{},
		})

		require.True(t, m.End())
	})

	t.Run("common", func(t *testing.T) {
		pageModel, err := pgm.New(
			coty.AppNameUndefined,
			hp.Body{},
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.MakeHeight(10)),
		)
		require.Nil(t, err)
		ri := MakeRowIterator(pageModel)

		require.False(t, ri.End())
		assert.Equal(t, row.MakeRow(7, nil), ri.RowModel())
		ri.Next()
	})

	t.Run("simple", func(t *testing.T) {
		expected := []string{
			"    header",
			"You motherfucker,",
			"come on you little",
			"ass… fuck with me,",
			"eh? You fucking",
			"little asshole,",
			"dickhead cocksucker…",
			"You fuckin' come",
			"on, come fuck with",
			":",
		}

		ri := MakeRowIterator(mustNewPageModel(
			"header",
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthTab,
					hp.MakeRowChunk("You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…"),
				),
				hp.MakeRow(size.WidthTab,
					hp.MakeRowChunk("You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker!"),
				),
			}),
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.MakeHeight(10)),
		))

		for i, ex := range expected {
			require.Equal(t, ex, ri.RowModel().String(), fmt.Sprintf("iteration %v: strings are not equal", i))
			ri.Next()
		}
	})
}

func mustNewPageModel(appName coty.NameApp, pageBody hp.Body, size ts.TerminalSize) *pgm.PageModel {
	hm, _ := pgm.New(appName, pageBody, size)
	return hm
}
