package data

import (
	"github.com/terryhay/dolly/pkg/dollyconf"
	"sort"
)

func getSortedCommands(commands map[dollyconf.Command]bool) (res []string) {
	res = make([]string, 0, len(commands))
	for command := range commands {
		res = append(res, string(command))
	}
	sort.Strings(res)

	return res
}

func getSortedFlags(groupFlagNameMap map[dollyconf.Flag]bool) (res []string) {
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

func getSortedFlagsForDescription(flagDescriptions map[dollyconf.Flag]*dollyconf.FlagDescription) (res []string) {
	res = make([]string, 0, len(flagDescriptions))
	for flag := range flagDescriptions {
		res = append(res, string(flag))
	}
	sort.Strings(res)

	return res
}
