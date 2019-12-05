// Package main contains the modules to create a FizzBuzz using Map
package main

import (
	"fmt"
	"runtime"

	fizzbuzzchan "github.com/abhishek-inspopindia/goworkspace/fizzbuzz_chan/pkg_fizzbuzz_chan"
)

func main() {
	runtime.GOMAXPROCS(2)
	fizz := 3
	buzz := 5

	m := fizzbuzzchan.FizzBuzz(fizz, buzz)

	fmt.Println(m)
}
