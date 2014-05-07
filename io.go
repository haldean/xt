package main

import (
	"bytes"
	"io"
)

const READ_BUFFER_SIZE = 16

func NewBuffer(r io.Reader) (Buffer, error) {
	var b Buffer
	b.Lines = make([]Line, 0)

	// bytes are read into this from the reader
	rb := make([]byte, READ_BUFFER_SIZE)
	// bytes are queued here until we find EOL
	lb := make([]byte, 0, READ_BUFFER_SIZE)
	for {
		n, err := r.Read(rb)
		for i := 0; i < n; i++ {
			lb = append(lb, rb[i])
			if rb[i] == '\n' {
				l := Line(bytes.Runes(lb))
				b.Lines = append(b.Lines, l)
				lb = make([]byte, 0, READ_BUFFER_SIZE)
			}
		}
		if err != nil && err == io.EOF {
			lb = append(lb, '\n')
			l := Line(bytes.Runes(lb))
			b.Lines = append(b.Lines, l)
			lb = make([]byte, 0, READ_BUFFER_SIZE)
			break
		} else if err != nil {
			return b, err
		}
	}
	return b, nil
}
