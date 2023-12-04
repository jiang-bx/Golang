package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(b []byte) {
	last := len(b) - 1

	for i := 0; i < len(b)/2; i++ {
		b[i], b[last-i] = b[last-i], b[i]
	}
}

func reverseUtf8(b []byte) {
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		fmt.Printf("r = %c\n", r)
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
}

func main() {
	b := []byte("一 二 三")
	fmt.Printf("b 的长度 = %d\n", len(b))
	reverseUtf8(b)
	fmt.Printf("%s\n", b)

	c := '测'

	fmt.Println(c)
}
