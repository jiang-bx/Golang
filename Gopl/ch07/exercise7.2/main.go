package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	w, c := CountingWriter(ioutil.Discard)

	fmt.Fprintf(w, "%s", "Hello, World!\n")
	fmt.Println(*c)
}

type ByteCounter struct {
	w       io.Writer
	writeen int64
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	n, err := b.w.Write(p)
	b.writeen += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := ByteCounter{w, 0}
	// c.w.Writer 有 io.Writer 接口中的 Writer 实现
	return &c, &c.writeen
}
