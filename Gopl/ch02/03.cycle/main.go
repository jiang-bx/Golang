package main

import "fmt"

var global *int

func f() {
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}

func main() {
	// 对于 f 函数中的 x 变量, 是在 堆内存空间分配, 原因在于 global 指向 x 变量, 用 Go 的术语说, 这个局部变量从 f 中逃逸了
	// 对于 g 函数中的 y 变量, 编译器可以选择在栈上分配, 也可以选择在堆上分配(然后由 GC 回收这个变量的内存空间)

	var (
		i, j, k int
	)

	i, j, k = 2, 3, 5

	j, k = k, j

	fmt.Printf("i = %d, j = %d, k = %d\n", i, j, k)
}
