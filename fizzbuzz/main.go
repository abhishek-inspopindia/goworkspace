// Package main contains the modules to create a FizzBuzz
package main

import (
	"fmt"

	fizzbuzz "github.com/abhishek-inspopindia/goworkspace/fizzbuzz/fizzbuzz.go"
)

func main() {
	fizz := 3
	buzz := 5

	m := fizzbuzz.FizzBuzz(fizz, buzz)
	fmt.Println(m)
}
