package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/terryhay/dolly/argparser/parsed"
	"github.com/terryhay/dolly/examples/example/dolly"
	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"
	fmtd "github.com/terryhay/dolly/tools/fmt_decorator"
)

const (
	exitCodeSuccess osd.ExitCode = iota
	exitCodeErrorArgParse
	exitCodeErrorGetFlagSl
	exitCodeErrorGetFlagIl
	exitCodeErrorGetFlagFl
	exitCodeConvertInt64Error
	exitCodeConvertFloat64Error
)

func main() {
	decOS := osd.New()
	decOut := fmtd.New()

	pd, err := dolly.Parse(os.Args[1:])
	if err != nil {
		decOS.Exit(exitCodeErrorArgParse, err)
	}

	decOS.Exit(entity(pd, decOut))
}

func entity(pd *parsed.Result, decOut fmtd.FmtDecorator) (code osd.ExitCode, err error) {
	switch {
	case pd.GetCommandMainName() == dolly.NameCommandNameless:
		var (
			builder strings.Builder

			values []parsed.ArgValue
		)

		values, err = pd.FlagArgValues(dolly.NameFlagSl)
		if err != nil {
			return exitCodeErrorGetFlagSl, err
		}

		if len(values) > 0 {
			builder.WriteString(fmt.Sprintf("flag %s arguments:\n\t", dolly.NameFlagSl))

			for _, v := range values {
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%s ", v)))
			}
			builder.WriteString("\n")
		}

		values, err = pd.FlagArgValues(dolly.NameFlagIl)
		if err != nil {
			return exitCodeErrorGetFlagIl, err
		}

		if len(values) > 0 {
			builder.WriteString(fmt.Sprintf("flag %s arguments:\n\t", dolly.NameFlagIl))

			var int64Value int64
			for _, v := range values {
				int64Value, err = v.Int64()
				if err != nil {
					return exitCodeConvertInt64Error, err
				}
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%d ", int64Value)))
			}
			builder.WriteString("\n")
		}

		values, err = pd.FlagArgValues(dolly.NameFlagFl)
		if err != nil {
			return exitCodeErrorGetFlagFl, err
		}

		if len(values) > 0 {
			builder.WriteString(fmt.Sprintf("flag %s arguments:\n\t", dolly.NameFlagFl))

			var float64Value float64
			for _, v := range values {
				float64Value, err = v.Float64()
				if err != nil {
					return exitCodeConvertFloat64Error, err
				}
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%f ", float64Value)))
			}
			builder.WriteString("\n")
		}

		decOut.Println(builder.String())

	case pd.GetCommandMainName() == dolly.NameCommandPrint:
		decOut.Println(fmt.Sprintf("command: %s", pd.GetCommandMainName()))

	}

	return exitCodeSuccess, nil
}
