package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 将 conn 中的内容 发送到 标准输出中
	go mustCopy(os.Stdout, conn)

	// 接受输入
	sc := bufio.NewScanner(os.Stdin)

CLOSE:
	for sc.Scan() {
		args := strings.Fields(sc.Text())
		cmd := args[0]
		switch cmd {
		case "close":
			fmt.Fprintln(conn, sc.Text())
			break CLOSE

		case "ls", "cd", "get":
			fmt.Fprintln(conn, sc.Text())
		case "send":
			if len(args) < 2 {
				log.Println("not enough argument")
			} else {
				filename := args[1]
				data, err := os.ReadFile(filename)
				if err != nil {
					log.Printf("read file err: %v", err)
				}
				fmt.Fprintf(conn, "%s %d\n", sc.Text(), countLines(data))
				fmt.Fprintf(conn, "%s", data)
			}
		}
	}
}

func countLines(data []byte) int {
	c := 0
	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		c++
	}
	return c
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
