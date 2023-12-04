package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	a, err := html.Parse(NewReader("a<h1>Hello<h11>"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse err: %v", err)
		os.Exit(1)
	}

	fmt.Println(a)

}

type StringReader string

func (s *StringReader) Read(p []byte) (int, error) {
	copy(p, *s)
	return len(*s), io.EOF
}

func NewReader(s string) io.Reader {
	sr := StringReader(s)
	return &sr
}
