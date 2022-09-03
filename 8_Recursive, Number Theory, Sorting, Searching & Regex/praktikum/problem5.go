package main

import "fmt"

func FindMinAndMax(arr []int) string {
    min, indexMin, max, indexMax := 0, 0, 0, 0
    for index, value := range arr {
        if value < min {
            min = value
            indexMin = index
        }

        if value > max {
            max = value
            indexMax = index
        }
    }
    

    return fmt.Sprintf("min: %d, index: %d, max: %d, index: %d", min, indexMin, max, indexMax)
}

func main() {
    fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
    fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
    fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
    fmt.Println(FindMinAndMax([]int{-1, 5, 5, 4, 2, 18}))
    fmt.Println(FindMinAndMax([]int{-2 ,5, -7, 4, 7, -20}))
}