package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")

var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}

// 构建: go build main.go
// 运行: ./main a bc def
// ./main -s / a bc def