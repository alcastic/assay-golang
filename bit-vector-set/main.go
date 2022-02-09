package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Sets in Go are usually implemented as a map[T]bool, where T is the element type.
// however, a bit vector is ideal in domains such as dataflow analysis where
// set elements are small non-negative integers, sets have many elements and set
// operations like union and intersection are common.

// An InSet is a set of small non-negative integers.
// Its zero value represents an empty set.
type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] ^= (1 << bit)
	}
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	for index, word := range s.words {
		if word != 0 {
			for bit := 0; bit < 64; bit++ {
				if word&(1<<bit) != 0 {
					fmt.Fprintf(&buf, " %d", (index*64 + bit))
				}
			}
		}
	}
	return "{" + strings.TrimLeft(buf.String(), " ") + "}"
}

func main() {
	var is IntSet
	fmt.Println(is.Has(0) || is.Has(1) || is.Has(63) || is.Has(64) || is.Has(65) || is.Has(256)) // false

	is.Add(64)
	fmt.Println(is.Has(64)) // true

	is.Remove(64)
	fmt.Println(is.Has(64)) // false

	is.Add(0)
	is.Add(1)
	is.Add(63)
	is.Add(64)
	is.Add(65)
	is.Add(256)
	fmt.Println(is.String()) // {0 1 63 64 65 256}
}
