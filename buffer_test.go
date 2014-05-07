package main

import (
	"strings"
	"testing"
)

var test_data = `
this is the test data
札幌でまたガ
  スボンベ爆発事件`

func runesEqual(r1 []rune, s string) bool {
	r2 := []rune(s)
	if len(r1) != len(r2) {
		return false
	}
	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}

func checkLineMatch(t *testing.T, l Line, s string) {
	if !runesEqual([]rune(l), s + "\n") {
		t.Errorf("line doesn't match: got %v, expected %v", l, []rune(s))
	}
}

func newTestBuffer(t *testing.T) (Buffer, bool) {
	r := strings.NewReader(test_data)
	buf, err := NewBuffer(r)
	if err != nil {
		t.Errorf("couldn't read buffer: %v\n", err)
		return buf, false
	}
	return buf, true
}

func TestRead(t *testing.T) {
	buf, ok := newTestBuffer(t)
	if !ok {
		return
	}
	if len(buf.Lines) != 4 {
		t.Errorf("expected 4 lines, got %d", len(buf.Lines))
		return
	}
	checkLineMatch(t, buf.Lines[0], "")
	checkLineMatch(t, buf.Lines[1], "this is the test data")
	checkLineMatch(t, buf.Lines[2], "札幌でまたガ")
	checkLineMatch(t, buf.Lines[3], "  スボンベ爆発事件")
}

func TestMove(t *testing.T) {
	buf, ok := newTestBuffer(t)
	if !ok {
		return
	}
	w := NewWindow(buf)
	if len(w.Cursors) != 1 {
		t.Errorf("expected 1 cursor, got %d", len(w.Cursors))
		return
	}
	if w.Cursors[0].X != 0 || w.Cursors[0].Y != 0 {
		t.Errorf("expected cursor at 0, 0; got %v", w.Cursors[0])
		return
	}

	w.Move(0, 2, 2)
	if w.Cursors[0].X != 2 || w.Cursors[0].Y != 2 {
		t.Errorf("after moving, cursor should be at 2, 2; got %v", w.Cursors[0])
		return
	}

	w.Move(0, 10, 0)
	if w.Cursors[0].X != 6 || w.Cursors[0].Y != 2 {
		t.Errorf("after moving, cursor should be at 6, 2; got %v", w.Cursors[0])
		return
	}
}
