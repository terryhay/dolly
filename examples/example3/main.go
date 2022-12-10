package main

import (
	"fmt"
	"github.com/terryhay/dolly/examples/example3/dolly"
	"os"
)

func main() {
	_, err := dolly.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("example3 error: %s", err.Error())
		os.Exit(1) // todo: resolve this problem
	}

	os.Exit(0)
}
