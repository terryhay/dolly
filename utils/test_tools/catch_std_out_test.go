package test_tools

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCatchStdOut(t *testing.T) {
	t.Parallel()

	text := gofakeit.Name()

	out := CatchStdOut(func() {
		fmt.Print(text)
	})

	require.Equal(t, text, out)
}

func TestMustBeNoError(t *testing.T) {
	defer func() {
		err := recover()
		require.NotNil(t, err)
	}()

	mustBeNoError(fmt.Errorf("surprise motherfucker"))
}
