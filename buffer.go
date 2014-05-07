package main

import (
	"fmt"
	"io"
)

type Line []rune

type Buffer struct {
	Lines []Line
}

type Cursor struct {
	X, Y uint32
}

type Window struct {
	Buf     Buffer
	Cursors []Cursor
}

func (c *Cursor) FitToBuffer(b Buffer) {
	if len(b.Lines) == 0 {
		c.X = 0
		c.Y = 0
		return
	}

	ny := uint32(len(b.Lines))
	if c.Y < 0 {
		c.Y = 0
	}
	if c.Y >= ny {
		c.Y = ny - 1
	}

	nx := uint32(len(b.Lines[c.Y]))
	if c.X < 0 {
		c.X = 0
	}
	if c.X >= nx {
		c.X = nx - 1
	}
}

func (w Window) Move(c, dx, dy uint32) {
	w.Cursors[c].X += dx
	w.Cursors[c].Y += dy
	w.Cursors[c].FitToBuffer(w.Buf)
}

func (w Window) Print(write io.Writer) {
	for y, l := range w.Buf.Lines {
		if y == int(w.Cursors[0].Y) {
			x := int(w.Cursors[0].X)
			fmt.Fprintf(write, "%v_%v", string(l[:x]), string(l[x:]))
		} else {
			fmt.Fprintf(write, "%v", string(l))
		}
	}
}

func NewWindow(b Buffer) Window {
	var w Window
	w.Buf = b
	w.Cursors = make([]Cursor, 1)
	return w
}
