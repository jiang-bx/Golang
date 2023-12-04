package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	a := 1
	fmt.Print(a)
	fmt.Printf("a = %d\n", a)
	fmt.Println(a)
}

func TestFprint(t *testing.T) {
	a := 1
	fmt.Fprint(os.Stdout, a)
	fmt.Fprintf(os.Stdout, "a = %d\n", a)
	fmt.Fprintf(os.Stdout, "a = %d", a)
	fmt.Println("")
}

func TestSprint(t *testing.T) {
	a := 1
	a1 := fmt.Sprint(a)
	fmt.Println("a1: ", a1)

	a2 := fmt.Sprintf("a = %d", a)
	fmt.Println("a2: ", a2)

	a3 := fmt.Sprintf("a = %d", a)
	fmt.Println("a3: ", a3)
}
