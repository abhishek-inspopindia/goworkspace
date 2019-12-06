// Package main contains the modules to create a FizzBuzz using Map
package main

import (
	"fmt"

	fizzbuzz "github.com/abhishek-inspopindia/goworkspace/fizzbuzz"
)

func main() {
	fizz := 3
	buzz := 5

	m := fizzbuzz.FizzBuzz(fizz, buzz)
	fmt.Println(m)
}
