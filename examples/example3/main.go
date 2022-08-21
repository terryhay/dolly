package main

import (
	"fmt"
	"github.com/terryhay/dolly/examples/example3/dolly"
	"os"
)

func main() {
	_, err := dolly.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("example3 error: %v", err.Error())
		os.Exit(int(err.Code()))
	}

	os.Exit(0)
}
