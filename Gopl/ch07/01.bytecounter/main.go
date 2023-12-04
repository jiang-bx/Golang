package main

import "fmt"

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func Test1() {
	var c ByteCounter

	c.Write([]byte("hello"))

	fmt.Println(c)

	c = 0

	var d ByteCounter = 0

	var name = "Dolly"
	var name2 = "word"

	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
	fmt.Fprintf(&d, "hello, %s", name2)
	fmt.Println(d)
}

func main() {
	Test1()
}
