package main

import (
	"bytes"
	"fmt"
)

type Inset struct {
	words []uint64
}

func (s *Inset) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *Inset) Add(x int) {
	word, bit := x/64, uint(x%64)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

func (s *Inset) UnionWidth(t *Inset) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] |= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

func (s *Inset) String() string {
	var buf bytes.Buffer

	buf.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}

	buf.WriteByte('}')

	return buf.String()
}

func (s *Inset) Len() int {
	return len(s.words)
}

func (s *Inset) Remove(x int) {
	//
}

func (s *Inset) Clear() {
	s.words = []uint64{}
}

func (s *Inset) Copy() *Inset {
	var res = new(Inset)
	copy(s.words, res.words)
	return res
}

func (s *Inset) AddAll(sl ...int) {
	for _, v := range sl {
		s.Add(v)
	}
}

func (s *Inset) Each(fn func(item uint64, index int)) {
	for i, v := range s.words {
		fn(v, i)
	}
}

func Test1() {
	var x, y Inset

	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWidth(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))
}

func Test2() {
	var a Inset

	a.AddAll(1, 2, 3)
	a.AddAll(4, 555, 6666)
	fmt.Println(a.String())
	a.Clear()
	fmt.Println(a.String())
	a.AddAll(4, 555, 6666)
	fmt.Println(a.String())
	fmt.Println(a.Len())

	a.Each(func(item uint64, index int) {
		fmt.Println(item, "----", index)
	})

	b := a.Copy()
	fmt.Println(b.String())
}

func main() {
	// Test1()
	Test2()
}
