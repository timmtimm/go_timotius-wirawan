package main

import (
	"fmt"
	"strings"
)

func compare(string1, string2 string) string {
	if len(string1) > len(string2) {
		return compare(string2, string1)
	}

	if strings.Contains(string1, string2) {
		return string2
	}

	compareDepan := compare(string1, string2[1:])
	compareBelakang := compare(string1, string2[:len(string2)-1])

	if len(compareDepan) > len(compareBelakang) {
		return compareDepan
	} else {
		return compareBelakang
	}
}

func main() {
    fmt.Println(compare("AKA", "AKASHI"))
	fmt.Println(compare("KANGOORO", "KANG"))
	fmt.Println(compare("KI", "KIJANG"))
	fmt.Println(compare("KUPU-KUPU", "KUPU"))
	fmt.Println(compare("ILALANG", "ILA"))
}