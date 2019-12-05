// Package main contains the modules to create a FizzBuzz
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fizz := 3
	buzz := 5

	m := make(map[int]string)

	FizzBuzz(fizz, buzz, m)

	fmt.Println(m)
}

// FizzBuzz - generate Key-Value pairs of the Map
func FizzBuzz(fizz, buzz int, m map[int]string) {
	for i := 1; i <= 100; i++ {
		if i%(fizz*buzz) == 0 {
			m[i] = "FizzBuzz"
		} else if i%fizz == 0 {
			m[i] = "Fizz"
		} else if i%buzz == 0 {
			m[i] = "Buzz"
		} else {
			m[i] = strconv.Itoa(i)
		}
	}
}
