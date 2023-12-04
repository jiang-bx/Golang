package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 10*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v", *period)
	time.Sleep(*period)
	fmt.Println()
}

// 自定义命令行标记

// go build main.go
// ./main
// ./main -period 100ms
// ./main -period 2m30s