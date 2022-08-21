package termbox_decorator

import (
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/helpdisplay/runes"
	"testing"
	"time"
)

func TestTermBoxDecorator(t *testing.T) {
	t.Parallel()

	tbDecor := NewTermBoxDecorator()
	require.Nil(t, tbDecor.Init())

	defer tbDecor.Close()

	err := tbDecor.Clear()
	require.Nil(t, err)

	w, h := tbDecor.Size()
	_ = w
	_ = h

	tbDecor.SetCell(0, 0, runes.RuneDot, termbox.ColorDefault, termbox.ColorDefault)
	tbDecor.SetRune(1, 1, runes.RuneLwQ)
	err = tbDecor.Flush()
	require.Nil(t, err)

	ch := make(chan bool, 1)
	go func() {
		time.Sleep(4 * time.Second)

		ch <- false
	}()

	go func() {
		tbDecor.PollEvent()
		ch <- true
	}()

	var ev termbox.Event

	select {
	case <-ch:
		break
	}

	require.NotNil(t, ev)
}
