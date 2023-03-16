package command_line_argument

import (
	"sort"
	"strings"

	"github.com/brianvoe/gofakeit"
)

type typeRandStorage struct {
	cmdArg       Argument
	cmdArgSecond Argument
}

var randStorage = func() typeRandStorage {
	args := []string{
		strings.ToLower(gofakeit.Color()),
		strings.ToLower(gofakeit.Color()),
	}
	sort.Strings(args)

	return typeRandStorage{
		cmdArg:       Argument(args[0]),
		cmdArgSecond: Argument(args[1]),
	}
}()

// RandCmdArg returns random Argument
func RandCmdArg() Argument {
	return randStorage.cmdArg
}

// RandCmdArgSecond returns random Argument
func RandCmdArgSecond() Argument {
	return randStorage.cmdArgSecond
}
