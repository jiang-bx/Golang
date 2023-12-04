package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	n := 0
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		n++
		fmt.Printf("连接者%d\n", n)
		go HandleConn(conn)
	}
}

func HandleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
