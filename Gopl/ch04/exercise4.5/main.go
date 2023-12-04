package main

import "fmt"

func remove(s []string) []string {
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			// 将 s[i+1:] 里的元素复制到 s[i:] 上面去,
			// 由于底层引用的是同一个数组
			// 所以会产生覆盖
			// s[i:]    ==> ["b","b","c","c","d","d","e"]
			// s[i+1:]  ==> ["b","c","c","d","d","e"]
			// 所以 s[i:] 变为  ["b","c","c","d","d","e","e"]
			// 在新的 s 中把最后一个舍弃掉
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}

func main() {
	s := []string{"a", "b", "b", "c", "c", "d", "d", "e"}
	s = remove(s)
	fmt.Println(s)
}
