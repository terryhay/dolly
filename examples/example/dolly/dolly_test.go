package dolly

import (
	"testing"
)

func TestDolly(t *testing.T) {
	t.Parallel()

	/*t.Run("simple", func(t *testing.T) {
		res, err := Parse(nil)
		require.Nil(t, err)
		require.True(t, cmp.Equal(&parsed.ParsedData{CommandID: CommandIDNamelessCommand}, res))
	})

	t.Run("incorrect_argument", func(t *testing.T) {
		res, err := Parse([]string{gofakeit.Color()})
		require.Nil(t, res)
		require.NotNil(t, err)
	})

	// todo: uncomment and fix this shit please
	/*t.Run("print_help_info", func(t *testing.T) {
			out := test_tools.CatchStdOut(func() {
				res, err := Parse([]string{"-h"})
				require.Nil(t, res)
				require.Nil(t, err)
			})

			ok, msg := test_tools.CheckSpaces(out)
			require.True(t, ok, msg)

			require.Equal(t, `[1mNAME[0m
		[1mexample[0m – shows how parser generator works

	[1mSYNOPSIS[0m
		[1mexample[0m [[1m-fl[0m [4mstr[0m [4m...[0m] [[1m-il[0m [4mstr[0m [4m...[0m] [[1m-sl[0m [4mstr[0m [4m...[0m]
		[1mexample print[0m [[1m-checkargs[0m] [[1m-f[0m [4mstr[0m] [[1m-fl[0m [4mstr[0m [4m...[0m] [[1m-i[0m [4mstr[0m] [[1m-il[0m [4mstr[0m [4m...[0m] [[1m-s[0m [4mstr[0m] [[1m-sl[0m [4mstr[0m [4m...[0m]

	[1mDESCRIPTION[0m
		you can write more detailed description here

		and use several paragraphs

	The commands are as follows:
		[1m<empty>[0m	checks arguments types

		[1mprint[0m	print command line arguments with optional checking
	The flags are as follows:
		[1m-checkargs[0m
			do arguments checking

		[1m-f[0m	single float

		[1m-fl[0m	float list

		[1m-i[0m	int string

		[1m-il[0m	int list

		[1m-s[0m	single string

		[1m-sl[0m	string list

	`,
				out)
		})//*/
}
