package plain_help_out

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"sort"
	"strings"
)

func getSortedCommands(commands map[apConf.Command]bool) (res []string) {
	if len(commands) == 0 {
		return nil
	}
	res = make([]string, 0, len(commands))
	for command := range commands {
		res = append(res, string(command))
	}
	sort.Strings(res)

	return res
}

func getSortedFlags(groupFlagNameMap map[apConf.Flag]bool) (res []string) {
	if len(groupFlagNameMap) == 0 {
		return nil
	}
	res = make([]string, 0, len(groupFlagNameMap))
	for flag := range groupFlagNameMap {
		res = append(res, string(flag))
	}
	sort.Strings(res)

	return res
}

func getSortedStrings(strings map[string]bool) (res []string) {
	if len(strings) == 0 {
		return nil
	}
	res = make([]string, 0, len(strings))
	for s := range strings {
		res = append(res, s)
	}
	sort.Strings(res)

	return res
}

type flagHelpOutData struct {
	NamePart        string
	DescriptionPart string
}

type byNamePart []flagHelpOutData

// Len implements Len sort interface method
func (i byNamePart) Len() int {
	return len(i)
}

// Less implements Less sort interface method
func (i byNamePart) Less(left, right int) bool {
	return i[left].NamePart < i[right].NamePart
}

// Swap implements Swap sort interface method
func (i byNamePart) Swap(left, right int) {
	i[left], i[right] = i[right], i[left]
}

func getSortedFlagsForDescription(flagDescriptions []*apConf.FlagDescription) (res []flagHelpOutData) {
	res = make([]flagHelpOutData, 0, len(flagDescriptions))
	for _, desc := range flagDescriptions {
		flags := make([]string, 0, len(desc.GetFlags()))
		for _, flag := range desc.GetFlags() {
			flags = append(flags, flag.ToString())
		}

		res = append(res, flagHelpOutData{
			NamePart:        strings.Join(flags, ", "),
			DescriptionPart: desc.GetDescriptionHelpInfo(),
		})

	}
	sort.Sort(byNamePart(res))

	return res
}
