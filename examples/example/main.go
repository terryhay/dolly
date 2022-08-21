package main

import (
	"fmt"
	"github.com/terryhay/dolly/examples/example/dolly"
	"github.com/terryhay/dolly/internal/os_decorator"
	"github.com/terryhay/dolly/pkg/parsed_data"
	"os"
	"strings"
)

const (
	exitCodeSuccess uint = iota
	exitCodeConvertInt64Error
	exitCodeConvertFloat64Error
)

func main() {
	osd := os_decorator.NewOSDecorator()

	pd, err := dolly.Parse(osd.GetArgs())
	if err != nil {
		fmt.Printf("example.Argparser error: %v\n", err.Error())
		os.Exit(int(err.Code()))
	}

	osd.Exit(logic(pd))
}

func logic(pd *parsed_data.ParsedData) (error, uint) {
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
			builder.WriteString(fmt.Sprintf("flag %v arguments:\n\t", dolly.FlagSl))

			for i = range values {
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%s ", values[i].ToString())))
			}
			builder.WriteString("\n")
		}

		if values, contain = pd.GetFlagArgValues(dolly.FlagIl); contain {
			builder.WriteString(fmt.Sprintf("flag %v arguments:\n\t", dolly.FlagIl))

			for i = range values {
				int64Value, err = values[i].ToInt64()
				if err != nil {
					return err, exitCodeConvertInt64Error
				}
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%d ", int64Value)))
			}
			builder.WriteString("\n")
		}

		if values, contain = pd.GetFlagArgValues(dolly.FlagFl); contain {
			builder.WriteString(fmt.Sprintf("flag %v arguments:\n\t", dolly.FlagFl))

			for i = range values {
				float64Value, err = values[i].ToFloat64()
				if err != nil {
					return err, exitCodeConvertFloat64Error
				}
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%f ", float64Value)))
			}
			builder.WriteString("\n")
		}

		fmt.Printf(builder.String())
	}

	return nil, exitCodeSuccess
}
