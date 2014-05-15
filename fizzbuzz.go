package main

import "fmt"

func main() {
	const (
		FIZZ     = "Fizz"
		BUZZ     = "Buzz"
		FIZZBUZZ = "FizzBuzz"
	)
	for i := 1; i <= 100; i++ {
		var s string
		if (i%3 == 0) && (i%5 == 0) {
			s = FIZZBUZZ
		} else if i%3 == 0 {
			s = FIZZ
		} else if i%5 == 0 {
			s = BUZZ
		}
		fmt.Printf("%3d - %s\n", i, s)
	}
}
