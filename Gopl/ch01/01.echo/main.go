package main

import (
	"fmt"
	"os"
)

func main() {
	// var s, sep string

	// sep = "   "

	// for i := 0; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// }
	// fmt.Println(s)

	s, sep := "", " "

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	}

	fmt.Println(s)

	fmt.Println(os.Args[1:])
}
