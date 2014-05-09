package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

type Rope struct {
	Left        *Rope
	Right       *Rope
	Weight      uint
	WeightBytes uint
	Data        string
}

var OutOfBounds error

const BALANCE_STRING_MIN = 16
const BALANCE_STRING_MAX = 64

func NewRope(s string) Rope {
	var r Rope
	r.Data = s
	r.rebalance()
	r.recalcWeights()
	return r
}

// make sure you recalculate weights after rebalancing!
func (r *Rope) rebalance() {
	if r.Left == nil && r.Right == nil {
		if utf8.RuneCountInString(r.Data) > BALANCE_STRING_MAX {
			c := len(r.Data) / 2
			for !utf8.ValidString(r.Data[:c]) && c < len(r.Data) {
				c += 1
			}
			var left, right Rope
			left.Data = r.Data[:c]
			right.Data = r.Data[c:]
			r.Left = &left
			r.Right = &right
			r.Data = ""
		}
	}
	if r.Left != nil {
		r.Left.rebalance()
	}
	if r.Right != nil {
		r.Right.rebalance()
	}
}

// returns (rune weight, byte weight)
func (r Rope) totalSubtreeWeights() (uint, uint) {
	rw := r.Weight
	bw := r.WeightBytes
	if r.Right != nil {
		rrw, rbw := r.Right.totalSubtreeWeights()
		rw += rrw
		bw += rbw
	}
	return rw, bw
}

func (r *Rope) recalcWeights() {
	if r.Left != nil {
		r.Left.recalcWeights()
	}
	if r.Right != nil {
		r.Right.recalcWeights()
	}
	r.Weight = uint(utf8.RuneCountInString(r.Data))
	r.WeightBytes = uint(len(r.Data))
	if r.Left != nil {
		rw, bw := r.Left.totalSubtreeWeights()
		r.Weight += rw
		r.WeightBytes += bw
	}
}

func (r Rope) Length() uint {
	l := r.Weight
	if r.Right != nil {
		l += r.Right.Length()
	}
	return l
}

func (r Rope) Bytes() uint {
	l := r.WeightBytes
	if r.Right != nil {
		l += r.Right.Bytes()
	}
	return l
}

func (r Rope) Assemble() string {
	var buf bytes.Buffer
	if r.Left != nil {
		buf.WriteString(r.Left.Assemble())
	}
	buf.WriteString(r.Data)
	if r.Right != nil {
		buf.WriteString(r.Right.Assemble())
	}
	return buf.String()
}

func (r Rope) Index(i uint) (rune, error) {
	if i < r.Weight {
		if r.Left != nil {
			return r.Left.Index(i)
		} else {
			for idx, c := range r.Data {
				if uint(idx) == i {
					return c, nil
				}
			}
			return utf8.RuneError, OutOfBounds
		}
	} else if r.Right != nil {
		return r.Right.Index(i - r.Weight)
	} else {
		return utf8.RuneError, OutOfBounds
	}
}

func find(r Rope, c rune, start uint, offset uint) int {
	if start < r.Weight {
		if r.Left != nil {
			i := find(*r.Left, c, start, 0)
			if i != -1 {
				return i + int(offset)
			}
		} else {
			i := strings.IndexRune(r.Data[start:], c)
			if i != -1 {
				return i + int(offset) + int(start)
			}
		}
	}
	if r.Right != nil {
		if start > r.Weight {
			start -= r.Weight
		} else {
			start = 0
		}
		return find(*r.Right, c, start, r.Weight)
	}
	return -1
}

func (r Rope) Find(c rune, start uint) int {
	return find(r, c, start, 0)
}

func (r Rope) debug(indent int) {
	fmt.Printf("\n")
	for i := 0; i < indent; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("rope runes=%d bytes=%d data=\"%s\"",
		r.Weight, r.WeightBytes, string(r.Data))
	if r.Left != nil {
		r.Left.debug(indent + 1)
	} else {
		fmt.Printf(" no-left")
	}
	if r.Right != nil {
		r.Right.debug(indent + 1)
	} else {
		fmt.Printf(" no-right")
	}
}

func (r Rope) Debug() {
	r.debug(0)
	fmt.Println()
}
