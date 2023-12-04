package main

import "fmt"

func rotate(s []int, n int) {
	n %= len(s)

	temp := append(s, s[:n]...)

	fmt.Println(temp)

	copy(s, temp[n:])
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotate(a, 3)
	fmt.Println(a)
}
