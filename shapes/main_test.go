package shapes

import (
	"fmt"
	"testing"
)

// Reference: https://dev.to/quii/learn-go-by-writing-tests-structs-methods-interfaces--table-driven-tests-1p01
func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Square", shape: Square{Side: 10}, hasArea: 100.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Square", shape: Square{Side: 10}, hasPerimeter: 40.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasPerimeter: 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}

func BenchmarkArea(b *testing.B) {

	areaBenchmarks := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Square", shape: Square{Side: 10}, hasArea: 100.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
	}

	for i := 0; i < b.N; i++ {
		for _, bb := range areaBenchmarks {
			// using bb.name from the case to use it as the `b.Run` test name
			b.Run(bb.name, func(b *testing.B) {
				_ = bb.shape.Area()
			})
		}
	}
}

func BenchmarkPerimeter(b *testing.B) {

	perimeterBenchmarks := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Square", shape: Square{Side: 10}, hasPerimeter: 40.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasPerimeter: 62.83185307179586},
	}

	for i := 0; i < b.N; i++ {
		for _, bb := range perimeterBenchmarks {
			// using bb.name from the case to use it as the `b.Run` test name
			b.Run(bb.name, func(b *testing.B) {
				_ = bb.shape.Perimeter()
			})
		}
	}
}

func ExampleShape_Area() {
	areaExamples := []struct {
		name  string
		shape Shape
	}{
		{name: "Square", shape: Square{Side: 10}},
		{name: "Circle", shape: Circle{Radius: 10}},
	}

	for _, ee := range areaExamples {
		fmt.Println(ee.shape.Area())
	}
	// Output:
	// 100
	// 314.1592653589793
}

func ExampleShape_Perimeter() {
	perimeterExamples := []struct {
		name  string
		shape Shape
	}{
		{name: "Square", shape: Square{Side: 10}},
		{name: "Circle", shape: Circle{Radius: 10}},
	}

	for _, ee := range perimeterExamples {
		fmt.Println(ee.shape.Perimeter())
	}
	// Output:
	// 40
	// 62.83185307179586
}
