package main

import "fmt"

func main() {
	TestScan()
}

func TestScan() {
	a := ""
	b := 0
	fmt.Println("请输入 a 和 b")
	fmt.Scan(&a, &b)
	fmt.Printf("您输入 a = %s, b = %d", a, b)
}
