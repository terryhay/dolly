package main

import (
	"fmt"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/examples/example/dolly"
	"os"
	"strings"
)

const (
	exitCodeSuccess uint = iota
	exitCodeConvertInt64Error
	exitCodeConvertFloat64Error
)

func main() {
	pd, err := dolly.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("example.Argparser error: %s\n", err.Error())
		os.Exit(1)
	}

	code, _ := logic(pd)
	os.Exit(int(code))
}

func logic(pd *parsed_data.ParsedData) (uint, error) {
	switch pd.GetCommandID() {
	case dolly.CommandIDNamelessCommand:
		var (
			builder      strings.Builder
			contain      bool
			err          error
			float64Value float64
			i            int
			int64Value   int64
			values       []parsed_data.ArgValue
		)

		if values, contain = pd.GetFlagArgValues(dolly.FlagSl); contain {
			builder.WriteString(fmt.Sprintf("flag %s arguments:\n\t", dolly.FlagSl))

			for i = range values {
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%s ", values[i].ToString())))
			}
			builder.WriteString("\n")
		}

		if values, contain = pd.GetFlagArgValues(dolly.FlagIl); contain {
			builder.WriteString(fmt.Sprintf("flag %s arguments:\n\t", dolly.FlagIl))

			for i = range values {
				int64Value, err = values[i].ToInt64()
				if err != nil {
					return exitCodeConvertInt64Error, err
				}
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%d ", int64Value)))
			}
			builder.WriteString("\n")
		}

		if values, contain = pd.GetFlagArgValues(dolly.FlagFl); contain {
			builder.WriteString(fmt.Sprintf("flag %s arguments:\n\t", dolly.FlagFl))

			for i = range values {
				float64Value, err = values[i].ToFloat64()
				if err != nil {
					return exitCodeConvertFloat64Error, err
				}
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%f ", float64Value)))
			}
			builder.WriteString("\n")
		}

		fmt.Printf(builder.String())
	}

	return exitCodeSuccess, nil
}
