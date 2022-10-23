package page

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"sort"
)

func getSortedCommands(commands map[apConf.Command]bool) (res []string) {
	res = make([]string, 0, len(commands))
	for command := range commands {
		res = append(res, string(command))
	}
	sort.Strings(res)

	return res
}

func getSortedFlags(groupFlagNameMap map[apConf.Flag]bool) (res []string) {
	res = make([]string, 0, len(groupFlagNameMap))
	for flag := range groupFlagNameMap {
		res = append(res, string(flag))
	}
	sort.Strings(res)

	return res
}

func getSortedStrings(strings map[string]bool) (res []string) {
	res = make([]string, 0, len(strings))
	for s := range strings {
		res = append(res, s)
	}
	sort.Strings(res)

	return res
}

func getSortedFlagsForDescription(flagDescriptions map[apConf.Flag]*apConf.FlagDescription) (res []string) {
	res = make([]string, 0, len(flagDescriptions))
	for flag := range flagDescriptions {
		res = append(res, string(flag))
	}
	sort.Strings(res)

	return res
}
