package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	if number == 2 { return true }
	if number <= 1 || number % 2 == 0 { return false }

	for i := 3; i <= int(math.Sqrt(float64(number))); i += 2 {
		if number % i == 0 { return false }
	}

	return true
}

func primeX(number int) int {
	position := 0

	i := 2
	for true {
		if primeNumber(i) {
			position++
			if position == number { return i }
		}

		i++
	}

	return 0
}

func main() {
    fmt.Println(primeX(1))
    fmt.Println(primeX(5))
    fmt.Println(primeX(8))
    fmt.Println(primeX(9))
    fmt.Println(primeX(10))
}