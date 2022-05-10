package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Object interface {
	volume() float64
}

type Geometry interface {
	Shape
	Object
	getColor() string
}

type Cube struct {
	edge  float64
	color string
}

func (c Cube) area() float64 {
	return 6 * (c.edge * c.edge)
}

func (c Cube) getColor() string {
	return c.color
}

func (c Cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

func measure(g Geometry) (float64, float64) {
	a := g.area()
	v := g.volume()

	return a, v
}

func main() {
	c := Cube{edge: 2}
	a, v := measure(c)
	fmt.Printf("Area %v, Volume %v\n", a, v)
}
