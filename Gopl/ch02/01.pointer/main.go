package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

var p = f()

func main() {
	*p = 44
	fmt.Printf("%p\n", p)
	fmt.Printf("%d\n", *p)
	fmt.Printf("%p\n", f())
	fmt.Println(f() == f())
}
