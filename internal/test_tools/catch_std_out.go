package test_tools

import (
	"io/ioutil"
	"os"
)

// CatchStdOut returns output to "os.Stdout" from "runnable" as string
func CatchStdOut(runnable func()) string {
	realStdout := os.Stdout
	defer func() {
		os.Stdout = realStdout
	}()

	r, fakeStdout, err := os.Pipe()
	mustBeNoError(err)
	os.Stdout = fakeStdout

	runnable()

	// Need to close here, otherwise ReadAll never gets "EOF".
	mustBeNoError(fakeStdout.Close())
	newOutBytes, err := ioutil.ReadAll(r)
	mustBeNoError(err)
	mustBeNoError(r.Close())

	return string(newOutBytes)
}

func mustBeNoError(err error) {
	if err != nil {
		panic(err)
	}
}
