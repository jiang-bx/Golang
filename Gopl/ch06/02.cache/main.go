package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

func test1() {
	mapping["ss"] = "sss"
	mapping["aa"] = "aaa"

	fmt.Println(Lookup("ss"))
	fmt.Println(Lookup("aa"))
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup1(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func test2() {
	cache.mapping["ss"] = "sss"
	cache.mapping["aa"] = "aaa"

	fmt.Println(Lookup1("ss"))
	fmt.Println(Lookup1("aa"))
}

func main() {
	test1()
	test2()
}
