package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func Test() {
	var w io.Writer
	fmt.Printf("%T\n", w)

	w = os.Stdout
	fmt.Printf("%T\n", w)

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)

	var x interface{} = time.Now()
	fmt.Printf("%T\n", x)
}



func main() {
	Test()
}
