package main

/**
太麻烦了
type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}
*/

/**
// 优化
type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}
*/

type Point struct {
	X, Y int
}

type Circle struct {
	Point  // 匿名成员
	Radius int
}

type Wheel struct {
	Circle // 匿名成员
	Spokes int
}

func main() {
	var w Wheel

	// 这样访问也很麻烦
	// w.Circle.Center.X = 8
	// w.Circle.Center.Y = 9
	// w.Circle.Radius = 5
	// w.Spokes = 20

	// go 特性:
	// 在声明时只声明一个成员对应数据类型, 而不是指成员名字, 这类成员就叫做匿名成员
	// 匿名成员的数据类型必须是命名的类型或者指向一个命名类型的指针
	// 匿名成员可以直接访问叶子属性, 减少代码量
	// 匿名成员并也是有一个隐式的名称的

	w.X = 1
	w.Y = 2
	w.Radius = 3
	w.Spokes = 33
}
