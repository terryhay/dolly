package main

import (
	"fmt"
	"github.com/terryhay/dolly/examples/example2/dolly"
	"os"
)

func main() {
	parsedData, err := dolly.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("example.Argparser error: %s", err.Error())
		os.Exit(1) // todo: resolve this problem
	}

	switch parsedData.GetCommandID() {
	case dolly.CommandIDNamelessCommand:
		fmt.Println("dir")
	}

	os.Exit(0)
}
