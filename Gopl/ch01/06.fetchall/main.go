package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)

	urls := []string{
		"https://baidu.com",
		"https://qq.com",
		"http://gopl.io",
		"https://godoc.org",
		"https://golang.org",
	}

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		// 从通道接收值
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
