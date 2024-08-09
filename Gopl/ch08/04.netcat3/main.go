package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		return
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Panicln("donw")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)

	conn.Close()

	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
