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

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "$%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{
		"shose": 50,
		"socks": 5,
	}

	mux := http.NewServeMux()

	// http.HandlerFunc 是一个类型
	// type HandlerFunc func(ResponseWriter, *Request).
	// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	// 	f(w, r)
	// }

	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))

	// http.ListenAndServe
	// func ListenAndServe(addr string, handler Handler) error {
	// 	server := &Server{Addr: addr, Handler: handler}
	// 	return server.ListenAndServe()
	// }

	// Handler
	// type Handler interface {
	// 	ServeHTTP(ResponseWriter, *Request)
	// }

	// mux 实现了 ServeHTTP 接口

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
