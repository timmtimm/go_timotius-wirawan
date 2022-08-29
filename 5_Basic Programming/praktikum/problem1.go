package main

import "fmt"

func main() {
    T := 20.0
	r := 4.0

	const pi = 3.14

	L := 2 * pi * r * (r+T)
	fmt.Println(L)
}