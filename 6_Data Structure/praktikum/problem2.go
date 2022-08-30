package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	showOnce := []int{}
    lengthAngka := len(angka)

	for i := 0; i<lengthAngka; i++ {
		flagOnlyOne := true
		angkaString := string(angka[i])

		for index, checkAngka := range angka {
			checkAngkaString := string(checkAngka)

			if (index != i && angkaString == checkAngkaString) {
				flagOnlyOne = false
			}
		}

		if (flagOnlyOne == true) {
			angkaInt, err := strconv.Atoi(angkaString)

			if (err != nil) {
				panic(err)
			}
		
			showOnce = append(showOnce, angkaInt)
		}
	}

	return showOnce
}

func main() {
    fmt.Println(munculSekali("1234321"))
    fmt.Println(munculSekali("76523752"))
    fmt.Println(munculSekali("12345"))
    fmt.Println(munculSekali("1122334455"))
    fmt.Println(munculSekali("0872504"))
}