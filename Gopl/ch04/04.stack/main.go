package main

import "fmt"

func createStack() []int {
	return []int{}
}

func popStack(stack []int) (int, []int) {
	if len(stack) < 0 {
		return -1, stack
	}

	return stack[len(stack)-1], stack[0 : len(stack)-1]
}

func pushStack(stack []int, item int) []int {
	return append(stack, item)
}

func removeStack(stack []int, i int) []int {
	copy(stack[i:], stack[i+1:])
	return stack[:len(stack)-1]
}

func main() {
	item := 0
	stack := createStack()

	stack = pushStack(stack, 1)
	stack = pushStack(stack, 2)
	stack = pushStack(stack, 3)
	stack = pushStack(stack, 4)
	stack = pushStack(stack, 5)

	fmt.Printf("%d\n", stack)

	item, stack = popStack(stack)
	fmt.Printf("%d, item = %d\n", stack, item)

	item, stack = popStack(stack)
	fmt.Printf("%d, item = %d\n", stack, item)

	stack = removeStack(stack, 1)
	fmt.Printf("%d\n", stack)
}
