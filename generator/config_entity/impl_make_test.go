package config_entity

import (
	"testing"

	"github.com/stretchr/testify/require"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestPlaceholdersByNames(t *testing.T) {
	t.Parallel()

	placeholderEmptyOpt := &confYML.PlaceholderOpt{}

	tests := []struct {
		caseName string

		conf                   *confYML.ArgParserConfig
		expPlaceholdersByNames map[coty.NamePlaceholder]*confYML.Placeholder
	}{
		{
			caseName: "nil_placeholder",
			conf: confYML.NewArgParserConfig(&confYML.ArgParserConfigOpt{
				Placeholders: []*confYML.PlaceholderOpt{
					nil,
				},
				Commands: []*confYML.CommandOpt{
					{},
				},
			}),
		},
		{
			caseName: "empty_command",
			conf: confYML.NewArgParserConfig(&confYML.ArgParserConfigOpt{
				Placeholders: []*confYML.PlaceholderOpt{
					placeholderEmptyOpt,
				},
				Commands: []*confYML.CommandOpt{
					{
						UsingPlaceholders: []string{
							placeholderEmptyOpt.Name,
						},
					}},
			}),
			expPlaceholdersByNames: map[coty.NamePlaceholder]*confYML.Placeholder{
				coty.NamePlaceholderUndefined: confYML.NewPlaceholder(placeholderEmptyOpt),
			},
		},
		{
			caseName: "duplicate_placeholders",
			conf: confYML.NewArgParserConfig(&confYML.ArgParserConfigOpt{
				Placeholders: []*confYML.PlaceholderOpt{
					placeholderEmptyOpt,
				},
				Commands: []*confYML.CommandOpt{
					{
						UsingPlaceholders: []string{
							placeholderEmptyOpt.Name,
							placeholderEmptyOpt.Name,
						},
					}},
			}),
			expPlaceholdersByNames: map[coty.NamePlaceholder]*confYML.Placeholder{
				coty.NamePlaceholderUndefined: confYML.NewPlaceholder(placeholderEmptyOpt),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseName, func(t *testing.T) {
			placeholders := placeholdersByNames(tc.conf)
			require.Equal(t, tc.expPlaceholdersByNames, placeholders)
		})
	}
}
