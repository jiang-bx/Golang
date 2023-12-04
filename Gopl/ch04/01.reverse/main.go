package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse1(s *[10]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}

	reverse(a[:])

	fmt.Println(a)

	b := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	reverse1(&b)
	fmt.Println(b)
}
