package main

import "fmt"

func playWithAsterisk(n int) {
	for i := 1; i<=n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Printf("* ")
		}
		fmt.Println("")
	}
}

func main() {
    playWithAsterisk(5)
}