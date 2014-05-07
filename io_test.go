package main

import (
	"io"
	"os"
	"testing"

	"github.com/dustin/randbo"
)

type fakeLineReader struct {
	i int
	lim int
}

func (f *fakeLineReader) Read(p []byte) (n int, err error) {
	if (f.i >= f.lim) {
		return 0, io.EOF
	}
	max := len(p)
	if f.i + max > f.lim {
		max = f.lim - f.i
	}
	x := 33 % len(p)
	for i := 0; i < max; i++ {
		if (i == x) {
			p[i] = '\n'
		} else {
			p[i] = 'g'
		}
	}
	f.i += max
	return max, nil
}

func benchStream(b *testing.B, length int) {
	b.SetBytes(int64(length))
	for i := 0; i < b.N; i++ {
		r := fakeLineReader{lim: length}
		_, err := NewBuffer(&r)
		if err != nil {
			b.Fatalf("got error when reading: %v\n", err)
		}
	}
}

func BenchmarkStream_16(b *testing.B) {
	benchStream(b, 16)
}

func BenchmarkStream_1024(b *testing.B) {
	benchStream(b, 1024)
}

func BenchmarkStream_8192(b *testing.B) {
	benchStream(b, 8192)
}

func BenchmarkStream_102400(b *testing.B) {
	benchStream(b, 102400)
}

func createRandomDataFile(b *testing.B, length int) string {
	f, err := os.Create("/tmp/test")
	if err != nil {
		b.Fatalf("couldn't create file: %v\n", err)
	}
	rand := randbo.New()
	io.CopyN(f, rand, int64(length))
	f.Close()
	return "/tmp/test"
}

func benchFile(b *testing.B, length int) {
	path := createRandomDataFile(b, length)
	f, err := os.Open(path)
	if err != nil {
		b.Fatalf("couldn't open file: %v\n", err)
	}
	b.SetBytes(int64(length))
	for i := 0; i < b.N; i++ {
		f.Seek(0, 0)
		_, err := NewBuffer(f)
		if err != nil {
			b.Fatalf("got error when reading: %v\n", err)
		}
	}
}

func BenchmarkFile_16(b *testing.B) {
	benchFile(b, 16)
}

func BenchmarkFile_1024(b *testing.B) {
	benchFile(b, 1024)
}

func BenchmarkFile_8192(b *testing.B) {
	benchFile(b, 8192)
}

func BenchmarkFile_102400(b *testing.B) {
	benchFile(b, 102400)
}
