package main

import (
	"fmt"
	"github.com/terryhay/dolly/examples/example2/dolly"
	"os"
)

func main() {
	parsedData, err := dolly.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("example.Argparser error: %v", err.Error())
		os.Exit(int(err.Code()))
	}

	switch parsedData.GetCommandID() {
	case dolly.CommandIDNamelessCommand:
		fmt.Println("dir")
	}

	os.Exit(0)
}
