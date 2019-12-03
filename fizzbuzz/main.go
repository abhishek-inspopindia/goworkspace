// Package Main
package main

import (
	"fmt"
	"strconv"
)

func main() {

	fizz := 3
	buzz := 5
	fizzbuzz := fizz * buzz

	m := make(map[int]string)

	c := make(chan string)

	go func() {
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
	}()

	cnt := 1
	for n := range c {
		m[cnt] = n
		cnt++
		//fmt.Println(n)
	}

	defer fmt.Println(m)
}
