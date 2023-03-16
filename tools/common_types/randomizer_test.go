package common_types

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/tools/size"
)

func TestRandomizer(t *testing.T) {
	t.Parallel()

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		checks := [...]error{
			RandError(),
			RandErrorSecond(),
			RandErrorThird(),
			RandErrorFourth(),
			RandErrorFifth(),
		}

		for i := 0; i < len(checks); i++ {
			require.Error(t, checks[i])

			for j := i + 1; j < len(checks); j++ {
				require.NotErrorIs(t, checks[i], checks[j])
			}
		}
	})

	t.Run("NameApp", func(t *testing.T) {
		t.Parallel()

		require.True(t, len(RandNameApp()) > 0)
	})

	t.Run("InfoChapter", func(t *testing.T) {
		t.Parallel()

		require.True(t, len(RandInfoChapterName()) > 0)
		require.True(t, len(RandInfoChapterDescription()) > 0)
		require.True(t, len(RandInfoChapterDescriptionSecond()) > 0)
	})

	t.Run("ArgPlaceholderID", func(t *testing.T) {
		t.Parallel()

		require.True(t, RandIDPlaceholder() > 0)
		require.True(t, RandIDPlaceholderSecond() > 0)
		require.True(t, RandIDPlaceholder() < RandIDPlaceholderSecond())
		require.True(t, RandIDPlaceholderSecond() < RandIDPlaceholderThird())
	})

	t.Run("NameCommand", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, RandNameCommand().IsValid(false))
		require.NoError(t, RandNameCommandSecond().IsValid(false))
		require.NoError(t, RandNameCommandThird().IsValid(false))
		require.NoError(t, RandNameCommandFourth().IsValid(false))
		require.NoError(t, RandNameCommandFifth().IsValid(false))

		require.True(t, RandNameCommand() < RandNameCommandSecond())
		require.True(t, RandNameCommandSecond() < RandNameCommandThird())
		require.True(t, RandNameCommandThird() < RandNameCommandFourth())
		require.True(t, RandNameCommandFourth() < RandNameCommandFifth())

		require.True(t, RandNameCommandShort() < RandNameCommandShortSecond())
	})

	t.Run("NamePlaceholder", func(t *testing.T) {
		t.Parallel()

		require.True(t, len(RandNamePlaceholder()) > 0)
		require.True(t, len(RandNamePlaceholderSecond()) > 0)

		require.True(t, RandNamePlaceholder() < RandNamePlaceholderSecond())
	})

	t.Run("NameFlagOneLetter", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, RandNameFlagOneLetter().IsValid())
		require.Equal(t, 2, len(RandNameFlagOneLetter()))
	})

	t.Run("NameFlagShort", func(t *testing.T) {
		t.Parallel()

		checks := [...]NameFlag{
			RandNameFlagShort(),
			RandNameFlagShortSecond(),
			RandNameFlagShortThird(),
		}

		for i := 1; i < len(checks); i++ {
			prev, cur := checks[i-1], checks[i]

			require.NoError(t, prev.IsValid())
			require.NoError(t, cur.IsValid())

			require.True(t, prev < cur)
		}
	})

	t.Run("NameFlagLong", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, RandNameFlagLong().IsValid())
	})

	t.Run("limitString", func(t *testing.T) {
		t.Parallel()

		s := gofakeit.Color()
		require.Equal(t, s[:len(s)-1], limitString(s, size.MakeWidth(len(s)-1)))
		require.Equal(t, s, limitString(s, size.MakeWidth(len(s)+1)))
	})
}

func TestAbs(t *testing.T) {
	t.Parallel()

	require.Equal(t, 1, abs(1))
	require.Equal(t, 1, abs(-1))
}
