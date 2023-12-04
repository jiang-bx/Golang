package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// 编写一个函数，原地将一个 UTF-8编码的[]byte类型的slice中相邻的空格 替换成一个空格返回
func replace(b []byte) []byte {
	for i := 0; i < len(b); {
		first, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(first) {
			second, _ := utf8.DecodeRune(b[i+size:])
			if unicode.IsSpace(second) {
				copy(b[i:], b[i+size:])
				b = b[:len(b)-size]
				continue
			}
		}
		i += size
	}
	return b
}

func main() {
	b := []byte(" 哈     喽   w      f    ")
	replace(b)
	fmt.Println(b)
}
