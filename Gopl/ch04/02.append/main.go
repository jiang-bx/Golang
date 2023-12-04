package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int

	zLen := len(x) + 1

	if zLen <= cap(x) {
		z = x[:zLen]
	} else {
		fmt.Println("扩容", "  ", y)
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z, x)
	}

	z[len(x)] = y

	return z
}

func main() {
	a := make([]int, 0)
	a = appendInt(a, 1)
	a = appendInt(a, 2)
	a = appendInt(a, 3)
	a = appendInt(a, 4)
	a = appendInt(a, 5)
	a = appendInt(a, 6)
	a = appendInt(a, 7)

	fmt.Println(a)

}
