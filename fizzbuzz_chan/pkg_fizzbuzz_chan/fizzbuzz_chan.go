// Package fizzbuzzchan contains the modules to create a FizzBuzz using Channels and Map
package fizzbuzzchan

import (
	"strconv"
)

// global variable declaration
var c = make(chan string)
var m = make(map[int]string)

// FizzBuzz - uses Channels and Maps to create FizzBuzz
func FizzBuzz(fizz, buzz int) map[int]string {
	go createFizzBuzz(fizz, buzz, c)
	mapFizzBuzz(m, c)

	return m
}

// CreateFizzBuzz - creates FizzBuzz and insert values into the Channel
func createFizzBuzz(fizz, buzz int, c chan string) {
	fizzbuzz := fizz * buzz
	for i := 1; i <= 100; i++ {
		if i%fizzbuzz == 0 {
			c <- "FizzBuzz"
		} else if i%fizz == 0 {
			c <- "Fizz"
		} else if i%buzz == 0 {
			c <- "Buzz"
		} else {
			c <- strconv.Itoa(i)
		}
	}
	close(c)
}

// MapFizzBuzz - maps the channel values to the global map variable
func mapFizzBuzz(m map[int]string, c chan string) {
	cnt := 1
	for n := range c {
		m[cnt] = n
		cnt++
	}
}
