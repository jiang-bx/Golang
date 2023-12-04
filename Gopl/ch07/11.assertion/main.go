package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Test1() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Println(f)
}

func Test2() {
	var w io.Writer
	w = os.Stdout
	if f, ok := w.(*os.File); ok {
		fmt.Println(f)
	}

	if c, ok := w.(*bytes.Buffer); !ok {
		fmt.Println(c)
		fmt.Println(ok)
	}
}

func main() {
	Test2()
}
