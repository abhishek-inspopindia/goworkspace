package main

import (
	"fmt"

	fizzbuzz "github.com/abhishek-inspopindia/goworkspace/fizzbuzz/pkg_fizzbuzz"
)

func main() {
	fizz := 3
	buzz := 5

	m := make(map[int]string)

	fizzbuzz.FizzBuzz(fizz, buzz, m)

	fmt.Println(m)
}
