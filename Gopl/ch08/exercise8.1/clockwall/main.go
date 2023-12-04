package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type server struct {
	name    string
	addr    string
	message string
}

func main() {
	// go run main.go a=localhost:8000  b=localhost:8000

	fmt.Println(os.Args)
	servers := parse(os.Args[1:])
	for _, ser := range servers {
		conn, err := net.Dial("tcp", ser.addr)
		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()

		go func(ser *server) {
			sc := bufio.NewScanner(conn)
			for sc.Scan() {
				ser.message = sc.Text()
			}
		}(ser)
	}

	for {
		fmt.Printf("\n")
		for _, server := range servers {
			fmt.Printf("%s: %s\n", server.name, server.message)
		}
		fmt.Printf("-------------")
		time.Sleep(2 * time.Second)
	}
}

func parse(args []string) []*server {
	var servers []*server

	for _, arg := range args {
		s := strings.SplitN(arg, "=", 2)
		servers = append(servers, &server{
			name:    s[0],
			addr:    s[1],
			message: "",
		})
	}

	return servers
}
