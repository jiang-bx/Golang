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
			log.Fatal(err)
			continue
		}

		n++

		fmt.Printf("连接者 %d\n", n)

		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("11:22:33\n"))
		if err != nil {
			fmt.Println("111")
			return
		}
		time.Sleep(2 * time.Second)
	}
}
