// Package main contains the modules to create a FizzBuzz
package main

import (
	"fmt"
)

func main() {
	fizzbuzz()
}

func fizzbuzz() {
	fizz := 3
	buzz := 5

	m := fizzbuzz.FizzBuzz(fizz, buzz)
	fmt.Println(m)
}
