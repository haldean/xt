package main

func (w *Window) Replace(c uint8, r rune) {
	cur := w.Cursors[c]
	w.Buf.Lines[cur.Y][cur.X] = r
}
