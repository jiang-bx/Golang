package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lr := LimitReader(strings.NewReader("FF2345"), 1)

	// ReadAll 中会调用  lr 实现的 Read 方法
	b, err := io.ReadAll(lr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}

	fmt.Printf("%s\n", b)
}

type LimitedReader struct {
	// 底层读取函数
	underlyingReader io.Reader
	// 限制字节数
	remainBytes      int64
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.remainBytes <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > r.remainBytes {
		p = p[:r.remainBytes]
	}

	// 在这里调用 r.underlyingReader.Read() 是由外部传入的
	// 在这里例子中, 就是 strings.NewReader("FF2345") 返回的 Read 方法
	n, err = r.underlyingReader.Read(p)

	r.remainBytes -= int64(n)

	return
}

// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// 这个函数返回的 reader 是 LimitedReader.Read() 方法, 因为 io.Reader 只是接口,
// 而 LimitedReader.Read() 是具体实现的实例
// 而且参数 r 的 Read 是接收外部实现的 Read 方法
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
