package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// func printCircle(c Circle) {
// 	fmt.Println("Shape:", c)
// 	fmt.Println("Area:", c.area())
// 	fmt.Println("Permiter:", c.permiter())

// }

// func printRectangle(r Rectangle) {
// 	fmt.Println("Shape:", r)
// 	fmt.Println("Area:", r.area())
// 	fmt.Println("Permiter:", r.permiter())
// }

func print(s Shape) {
	fmt.Printf("Shape %#v\n", s)
	fmt.Printf("Area %v\n", s.area())
	fmt.Printf("Perimeter %v\n", s.perimeter())
}

func main() {
	// c1 := Circle{radius: 5.}
	// r1 := Rectangle{width: 3., height: 2.1}

	// print(c1)
	// print(r1)

	// Interface Dynamic type and Polymorphism

	// var s Shape
	// fmt.Printf("%T\n", s)

	// ball := Circle{radius: 2.5}
	// s = ball

	// print(s)
	// fmt.Printf("%T\n", s)

	// room := Rectangle{width: 2, height: 3}
	// s = room
	// fmt.Printf("%T\n", s)

	// Type assertions and type switches

	var s Shape = Circle{radius: 2.5}
	fmt.Printf("%T\n", s)

	ball, ok := s.(Circle)

	if ok == true {
		fmt.Printf("Ball volume %v\n", ball.volume())
	}

	s = Rectangle{width: 3.4, height: 2.2}
	switch value := s.(type) {
	case Circle:
		fmt.Printf("%#v has circle type\n", value)
	case Rectangle:
		fmt.Printf("%#v has rectangle type\n", value)
	}

}
