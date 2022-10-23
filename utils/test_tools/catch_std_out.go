package test_tools

import (
	"io"
	"os"
)

// CatchStdOut returns output to "os.Stdout" from "runnable" as string
func CatchStdOut(runnable func()) string {
	realStdout := os.Stdout
	defer func() {
		os.Stdout = realStdout
	}()

	r, fakeStdout, errPipe := os.Pipe()
	mustBeNoError(errPipe)
	os.Stdout = fakeStdout

	runnable()

	// Need to close here, otherwise ReadAll never gets "EOF".
	mustBeNoError(fakeStdout.Close())
	newOutBytes, errReadAll := io.ReadAll(r)
	mustBeNoError(errReadAll)
	mustBeNoError(r.Close())

	return string(newOutBytes)
}

func mustBeNoError(err error) {
	if err != nil {
		panic(err)
	}
}
