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

func primaSegiEmpat(high, wide, start int) {
    primeFound := []int{}
    position := 0
    countPrimeFound := 0

    i := 2
	for true {
		if primeNumber(i) {
			position++

            if position > start {
                primeFound = append(primeFound, i)
                countPrimeFound++
            }

			if countPrimeFound == high * wide { break }
		}

		i++
	}

    // fmt.Println(primeFound)

    for i := 0; i < high*wide; i++ {
		fmt.Printf("%d ", primeFound[i])
		if (i+1) % wide == 0 { fmt.Println("") }
    }
}

func main() {
    primaSegiEmpat(2, 3, 13)
    primaSegiEmpat(5, 2, 1)
}