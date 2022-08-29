package main

import "fmt"

func primeNumber(number int) bool {
	if (number == 1) {
		return false
	} else if (number > 1) {
		for i := 2; i < number; i++ {
			if (number % i ==0) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func main() {
    fmt.Println(primeNumber(11))
    fmt.Println(primeNumber(13))
    fmt.Println(primeNumber(17))
    fmt.Println(primeNumber(20))
    fmt.Println(primeNumber(35))
}