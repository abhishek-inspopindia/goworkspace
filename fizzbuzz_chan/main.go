// Package main contains the modules to create a FizzBuzz using Map
package main

import (
	"fmt"

	fizzbuzzchan "github.com/abhishek-inspopindia/goworkspace/fizzbuzz_chan/pkg_fizzbuzz_chan"
)

func main() {
	fizz := 3
	buzz := 5

	m := fizzbuzzchan.FizzBuzz(fizz, buzz)

	fmt.Println(m)
}
