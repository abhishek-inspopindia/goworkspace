// Package main contains the modules to create a FizzBuzz
package main

import (
	"fmt"
)

func main() {
	fizz := 3
	buzz := 5

	m := make(map[int]string)

	FizzBuzz(fizz, buzz, m)

	fmt.Println(m)
}
