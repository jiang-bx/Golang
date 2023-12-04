package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	s := "Hello, World!\nHello, 世界！"

	var wc WordCounter
	fmt.Fprintf(&wc, "%s\n", s)
	fmt.Println(wc)

	var lc LineCounter
	fmt.Fprintf(&lc, "%s\n", s)
	fmt.Println(lc)
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w++
	}

	return len(p), nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*l++
	}

	return len(p), nil
}
