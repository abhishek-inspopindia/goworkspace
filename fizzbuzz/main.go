package main

import (
	"fmt"

	"github.com/abhishek-inspopindia/goworkspace/pkg/fizzbuzz"
)

func main() {
	fizz := 3
	buzz := 5

	m := fizzbuzz.FizzBuzz(fizz, buzz)
	fmt.Println(m)
}
