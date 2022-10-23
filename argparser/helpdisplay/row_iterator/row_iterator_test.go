package row_iterator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/argparser/helpdisplay/page"
	pgm "github.com/terryhay/dolly/argparser/helpdisplay/page_model"
	"github.com/terryhay/dolly/argparser/helpdisplay/row"
	rllMock "github.com/terryhay/dolly/argparser/helpdisplay/row_len_limiter/mock"
	"github.com/terryhay/dolly/argparser/helpdisplay/size"
	ts "github.com/terryhay/dolly/argparser/helpdisplay/terminal_size"
	"testing"
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
		pageModel, err := pgm.NewPageModel(
			page.Page{},
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.Height(10)),
		)
		require.Nil(t, err)
		ri := MakeRowIterator(pageModel)

		require.False(t, ri.End())
		assert.Equal(t, row.MakeRow(7, nil), ri.Row())
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
			page.Page{
				Header: page.MakeParagraph(0, "header"),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(1, "You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…"),
					page.MakeParagraph(1, "You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker!"),
				},
			},
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.Height(10)),
		))

		for i, ex := range expected {
			require.Equal(t, ex, ri.Row().String(), fmt.Sprintf("iteration %v: strings are not equal", i))
			ri.Next()
		}
	})
}

func mustNewPageModel(pageData page.Page, size ts.TerminalSize) *pgm.PageModel {
	hm, _ := pgm.NewPageModel(pageData, size)
	return hm
}
