package main

import "fmt"

func palindrome(input string) bool {
	lengthInput := len(input)

	for i := 0; i<lengthInput/2; i++ {
		if (input[i] != input[lengthInput - i - 1]) {
			return false
		}
	}
	
	return true
}

func main() {
    fmt.Println(palindrome("civic"))
    fmt.Println(palindrome("katak"))
    fmt.Println(palindrome("kasur rusak"))
    fmt.Println(palindrome("mistar"))
    fmt.Println(palindrome("lion"))
}