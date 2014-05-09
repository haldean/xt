package main

import (
	"testing"
)

func makeTestRope() Rope {
	var r, r2, r3, r4, r5 Rope
	r.Left = &r2
	r.Right = &r3
	r2.Left = &r4
	r2.Right = &r5

	r4.Data = "someday mother "
	r5.Data = "will die and I'll get "
	// these chinese characters take 3 bytes each
	r3.Data = "the money 世界"
	r.recalcWeights()

	return r
}

func TestWeights(t *testing.T) {
	checkWeight := func(r Rope, rweight, bweight uint, name string) {
		if r.Weight != rweight {
			t.Errorf("rope %v has bad rune weight %d, expected %d", name, r.Weight,
				rweight)
		}
		if r.WeightBytes != bweight {
			t.Errorf("rope %v has bad byte weight %d, expected %d", name,
				r.WeightBytes, bweight)
		}
	}

	r := makeTestRope()
	checkWeight(*r.Left.Left, 15, 15, "r4")
	checkWeight(*r.Left.Right, 22, 22, "r5")
	checkWeight(*r.Left, 15, 15, "r2")
	checkWeight(*r.Right, 12, 16, "r3")
	checkWeight(r, 37, 37, "r")
}

func TestLength(t *testing.T) {
	r := makeTestRope()
	if r.Length() != 49 {
		t.Errorf("rope has bad length %d, expected 49", r.Length)
	}
	if r.Bytes() != 53 {
		t.Errorf("rope has bad byte count %d, expected 53", r.Bytes())
	}
}

func TestAssemble(t *testing.T) {
	r := makeTestRope()
	str := r.Assemble()
	expect := "someday mother will die and I'll get the money 世界"
	if str != expect {
		t.Errorf("bad assemble, got \"%s\", expected \"%s\"", str, expect)
	}
}

func TestIndex(t *testing.T) {
	checkIndex := func(r Rope, i uint, expect rune) {
		c, err := r.Index(i)
		if err != nil {
			t.Errorf("got error when indexing at %d: %v\n", i, err)
		}
		if c != expect {
			t.Errorf("bad rune at index %d, got '%c', expected '%c'", i, c,
				expect)
		}
	}
	r := makeTestRope()
	checkIndex(r, 0, 's')
	checkIndex(r, 17, 'l')
	checkIndex(r, 47, '世')
}

func TestFind(t *testing.T) {
	checkFind := func(r Rope, c rune, start uint, expect int) {
		i := r.Find(c, start)
		if i != expect {
			t.Errorf("bad result for find %c, got %d, expected %d", c, i,
				expect)
		}
	}
	r := makeTestRope()
	checkFind(r, 's', 0, 0)
	checkFind(r, 's', 1, -1)
	checkFind(r, 'i', 10, 16)
	checkFind(r, 'i', 17, 21)
}

func TestNewRope(t *testing.T) {
	str := `
	why is the world in love again.
	why are we marching hand in hand.
	why are the ocean levels rising up.
	it's a brand new record
	for nineteen ninety
	they might be giants
	brand new album
	floooooooooooood`

	r := NewRope(str)
	a := r.Assemble()
	if a != str {
		t.Errorf("new/assemble was lossy, got \"%s\", expected \"%s\"", a, str)
	}
}

func TestConcat(t *testing.T) {
	str := "ana ng and I are getting old but we still haven't walked in the glow of each other's majestic presence"
	r1 := NewRope("ana ng and I are getting old but we still haven't ")
	r2 := NewRope("walked in the glow of each other's majestic presence")
	r := Concat(r1, r2)
	rlen := r.Length()
	elen := r1.Length() + r2.Length()
	if rlen != elen {
		t.Errorf("result of concat has wrong length %d, expected %d",
		rlen, elen)
	}
	a := r.Assemble()
	if a != str {
		t.Errorf("bad concat, got \"%s\", expected \"%s\"", a, str)
	}
}
