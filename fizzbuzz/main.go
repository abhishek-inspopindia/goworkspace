// Package Main
package main

import (
	"fmt"
	"strconv"
)

// global variable declaration
var c = make(chan string)
var m = make(map[int]string)

func main() {
	fizz := 3
	buzz := 5

	go CreateFizzBuzz(fizz, buzz)

	MapFizzBuzz()

	fmt.Println(m)
}

// CreateFizzBuzz - uses Channels to generate Key-Value pairs of the Map
func CreateFizzBuzz(fizz, buzz int) {
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
func MapFizzBuzz() {
	cnt := 1
	for n := range c {
		m[cnt] = n
		cnt++
	}
}
