// Package fizzbuzz contains the modules to create a FizzBuzz
package fizzbuzz

import (
	"strconv"
)

// FizzBuzz - generate Key-Value pairs of the Map
func FizzBuzz(fizz, buzz int) map[int]string {
	m := make(map[int]string)

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

	return m
}
