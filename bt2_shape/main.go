package main

import (
	"fmt"
	"math"
)

/*
Bài tập 2: Giao diện hình học
Tạo một interface Shape với các phương thức:
Area() float64
Perimeter() float64
Tạo các struct Circle, Rectangle và Triangle, và triển khai interface Shape cho từng struct.
Viết hàm nhận một slice các Shape và trả về tổng diện tích và chu vi của tất cả các hình.
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
func (t Triangle) Perimeter() float64 {
	return t.base + t.height
}

func sumArea(shapes []Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func sumPerimeter(shapes []Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.Perimeter()
	}
	return perimeter
}
func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 6}
	triangle := Triangle{base: 7, height: 3}

	shapes := []Shape{circle, rectangle, triangle}
	fmt.Println("Sum area:", sumArea(shapes))
	fmt.Println("Sum perimeter:", sumPerimeter(shapes))
}
