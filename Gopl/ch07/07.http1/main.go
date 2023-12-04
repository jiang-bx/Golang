package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		// price 会调用 toString 方法, 也就是 dollars 的 toString
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{
		"shose": 50,
		"socks": 5,
	}

	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
