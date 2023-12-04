package main

import "fmt"

func Nonempty(strArr []string) []string {
	i := 0
	for _, s := range strArr {
		if s != "" {
			strArr[i] = s
			i++
		}
	}

	return strArr[:i]
}

func Nonempty2(strArr []string) []string {
	out := strArr[:0]

	for _, s := range strArr {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func test1() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", Nonempty(data))
	fmt.Printf("%q\n", data)
}

func test2() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", Nonempty2(data))
	fmt.Printf("%q\n", data)
}

func main() {
	test1()
	test2()
}
