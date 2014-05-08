package main

import (
	"testing"
)

func TestReplace(t *testing.T) {
	buf, ok := newTestBuffer(t)
	if !ok {
		return
	}
	w := NewWindow(buf)
	if len(w.Cursors) != 1 {
		t.Errorf("expected 1 cursor, got %d", len(w.Cursors))
		return
	}

	w.Move(0, 1, 1)
	w.Replace(0, 'x')
	checkLineMatch(t, w.Buf.Lines[1], "txis is the test data")
}
