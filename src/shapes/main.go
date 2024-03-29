package main

import (
	"fmt"

	shapes "github.com/abhishek-inspopindia/goworkspace/pkg/shapes"
)

func main() {
	s := shapes.Square{Side: 10}
	c := shapes.Circle{Radius: 10}

	fmt.Printf("Type(s) = %T\tSide(s) = %v\tArea(s) = %v\tPerimeter(s) = %v\n", s, s.Side, s.Area(), s.Perimeter())
	fmt.Printf("Type(c) = %T\tRadius(c) = %v\tArea(c) = %v\tPerimeter(c) = %v\n", c, c.Radius, c.Area(), c.Perimeter())
}
