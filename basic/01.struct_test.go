package basic

import (
	"fmt"
	"testing"
)

func TestT1(t *testing.T) {
	type User struct {
		ac bool
		na string
	}

	u := User{
		ac: true,
		na: "fdsf",
	}

	// &u 类型为 *User
	fmt.Println(&u)

	fmt.Println((&u).na)

	fmt.Println(u)

	b := 10

	// &变量名, 类型为 *int
	c := &b

	fmt.Println(b)
	fmt.Println(c)
	// *变量名(类型为 *T 时), 取地址所对应的值
	fmt.Println(*c)

	// u1 为指针类型 *User
	u1 := &User{
		ac: false,
		na: "falsesss",
	}

	fmt.Println(u1)
	fmt.Println((*u1).na)

	// 这种方式是相同的, 取子属性的值
	fmt.Println(&(u1.ac) == &u1.ac)
}
