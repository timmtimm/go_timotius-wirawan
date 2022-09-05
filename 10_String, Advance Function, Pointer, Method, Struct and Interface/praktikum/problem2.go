package main

import "fmt"

func caesar(offset int, input string) string {
	var result string

	for i:=0; i<len(input); i++ {
		if input[i] > 'A' && input[i] < 'Z' { 
			result += string(int(int(input[i]) + offset - 65) % 26 + 65)
		} else {
			result += string(int(int(input[i]) + offset - 97) % 26 + 97)
		}
	}

	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}