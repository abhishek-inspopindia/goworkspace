//Package shapes provides various types of shapes and their informations
package shapes

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}
type Circle struct {
	radius float64
}
type Shape interface {
	area() float64
	circumference() float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) circumference() float64 {
	return 4 * s.side
}

func (c Circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	s := Square{10}
	c := Circle{10}

	fmt.Printf("Type(s) = %T\tSide(s) = %v\tArea(s) = %v\tCircumference(s) = %v\n", s, s.side, s.area(), s.circumference())
	fmt.Printf("Type(c) = %T\tRadius(c) = %v\tArea(c) = %v\tCircumference(c) = %v\n", c, c.radius, c.area(), c.circumference())
}
