package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// test1()
	// test2()
}

func Test3() {
	var any interface{}
	any = true
	any = 12.45
	any = "hello"
	any = new(bytes.Buffer)

	fmt.Println(any)
}

type Inset struct {
	a string
}

func (i *Inset) String() string {
	return i.a
}

func (i Inset) String1() string {
	return i.a
}

func Test2() {
	var i Inset = Inset{
		a: "123",
	}
	fmt.Println(i.String())
	fmt.Println(i.String1())

	var j *Inset = &Inset{
		a: "456",
	}
	fmt.Println(j.String())
	fmt.Println(j.String1())

}

func Test1() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	fmt.Println(w)

	var wrc io.ReadWriteCloser
	wrc = os.Stdout
	// wrc = new(bytes.Buffer) // missing method Close
	fmt.Println(wrc)

	w = wrc
	// wrc = w // missing method Close
}
