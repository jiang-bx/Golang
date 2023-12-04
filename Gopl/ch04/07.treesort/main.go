package main

import "fmt"

type tree struct {
	value int
	left  *tree
	right *tree
}

func sort(values []int) {
	var root *tree

	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func fTest(t *tree) *tree {
	a := &tree{
		value: 111,
	}

	return a
}

func main() {
	v := []int{4, 2, 66, 7, 23, 1, 5}
	sort(v)

	// 新的结构体
	var a = tree{
		value: 1,
	}

	// 新的结构体
	var b = tree{
		value: 2,
	}

	var c = fTest(&a)
	var d = fTest(&b)

	fmt.Println(c, d)
}
