package main

import "fmt"

func pangkat(base, pangkat int) int {
	var result int = 1

	for pangkat > 0 {
		if (pangkat % 2 == 1) {
			result = result * base
		}

		pangkat = pangkat >> 1
		base = base * base
	}

	return result
}

func main() {
    fmt.Println(pangkat(2, 3))
    fmt.Println(pangkat(5, 3))
    fmt.Println(pangkat(10, 2))
    fmt.Println(pangkat(2, 5))
    fmt.Println(pangkat(7, 3))
}