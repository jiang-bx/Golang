package main

import "fmt"

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, q.Y - q.Y}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func test1() {
	var arr = Path{
		Point{1, 2}, Point{2, 3},
	}

	fmt.Println(arr)

	arr.TranslateBy(Point{3, 4}, true)

	fmt.Println(arr)
}

func main() {
	test1()
}
