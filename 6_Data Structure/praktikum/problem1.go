package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	if(len(arrayA) == 0)  {
		return arrayB
	} else if(len(arrayB) == 0) {
		return arrayA
	} else if(len(arrayA) == 0 && len(arrayB) == 0) {
		return []string{}
	} else {
		mergedArray := arrayA

		for _, wordB := range arrayB {
			sameWord := false

			for _, wordA := range arrayA {
				if (wordB == wordA) {
					sameWord = true
					break
				}
			}

			if (sameWord == false) {
				mergedArray = append(mergedArray, wordB)
			}
		}

		return mergedArray
	}
}

func main() {
    fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
    fmt.Println(ArrayMerge([]string{"sergei","jin"}, []string{"jin", "steve", "bryan"}))
    fmt.Println(ArrayMerge([]string{"alisa","yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
    fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
    fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
    fmt.Println(ArrayMerge([]string{}, []string{}))
}