package arg_parser_config

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNamelessCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	namelessCommandDescription := NewNamelessCommandDescription(
		CommandID(gofakeit.Uint32()),
		gofakeit.Name(),
		&ArgumentsDescription{},
		map[Flag]bool{Flag(gofakeit.Name()): true},
		map[Flag]bool{Flag(gofakeit.Name()): true},
	)

	commandDescription := namelessCommandDescription.(*CommandDescription)

	require.Equal(t, commandDescription.ID, namelessCommandDescription.GetID())
	require.Equal(t, commandDescription.DescriptionHelpInfo, namelessCommandDescription.GetDescriptionHelpInfo())
	require.Equal(t, commandDescription.ArgDescription, namelessCommandDescription.GetArgDescription())
	require.Equal(t, commandDescription.RequiredFlags, namelessCommandDescription.GetRequiredFlags())
	require.Equal(t, commandDescription.OptionalFlags, namelessCommandDescription.GetOptionalFlags())
}
