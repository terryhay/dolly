package string_builder_tools

import (
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/tools/size"
)

func TestStringBuilderTools(t *testing.T) {
	t.Parallel()

	t.Run("append", func(t *testing.T) {
		t.Parallel()

		var builder *strings.Builder
		require.Nil(t, Append(builder))

		builder = &strings.Builder{}
		textChunks := []string{gofakeit.Name(), gofakeit.NameSuffix()}

		builder = Append(builder, textChunks...)
		require.Equal(t, strings.Join(textChunks, ""), builder.String())
	})

	t.Run("break_row", func(t *testing.T) {
		t.Parallel()

		var builder *strings.Builder
		require.Nil(t, BreakRow(builder))

		builder = &strings.Builder{}
		builder = BreakRow(builder)
		require.Equal(t, ChunkBreakRow, builder.String())
	})

	t.Run("new_row", func(t *testing.T) {
		t.Parallel()

		var builder *strings.Builder
		require.Nil(t, NewRow(builder, size.WidthZero))

		textChunks := []string{gofakeit.Name(), gofakeit.NameSuffix()}

		builder = &strings.Builder{}
		builder = NewRow(builder, size.WidthZero, textChunks...)
		require.Equal(t, strings.Join(append([]string{ChunkBreakRow}, textChunks...), ""), builder.String())
	})

	t.Run("new_tab_row", func(t *testing.T) {
		t.Parallel()

		var builder *strings.Builder
		require.Nil(t, NewRow(builder, size.WidthTab))

		textChunks := []string{gofakeit.Name(), gofakeit.NameSuffix()}

		builder = &strings.Builder{}
		builder = NewRow(builder, size.WidthTab, textChunks...)
		require.Equal(t, strings.Join(append([]string{ChunkBreakRow, strings.Repeat(ChunkSpace, size.WidthTab.Int())}, textChunks...), ""), builder.String())
	})
}
