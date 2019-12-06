//Package shapes provides various types of shapes and their informations
package shapes

import (
	"math"
)

type Square struct {
	Side float64
}
type Circle struct {
	Radius float64
}

// Shape defines an interface having methods for Area() and Circumference()
type Shape interface {
	Area() float64
	Perimeter() float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// func main() {
// 	s := Square{10}
// 	c := Circle{10}

// 	fmt.Printf("Type(s) = %T\tSide(s) = %v\tArea(s) = %v\tPerimeter(s) = %v\n", s, s.Side, s.Area(), s.Perimeter())
// 	fmt.Printf("Type(c) = %T\tRadius(c) = %v\tArea(c) = %v\tPerimeter(c) = %v\n", c, c.Radius, c.Area(), c.Perimeter())
// }
