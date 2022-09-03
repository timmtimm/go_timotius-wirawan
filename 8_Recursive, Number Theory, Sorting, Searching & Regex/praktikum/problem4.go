package main

import (
    "fmt"
)

func maxArraySum(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func MaxSequence(arr []int) int {
    if len(arr) < 1 {
        return 0
    }

    sum := make([]int, len(arr) +1)
    maxSum := 0

    for index, value := range arr {
        sum[index+1] = value + sum[index]
        maxSum = maxArraySum(maxSum, sum[index+1])
    }

    return (maxArraySum(maxSum, MaxSequence(arr[1:])))
}

func main() {
    // fmt.Println(compareValue(2, 1))
    fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, -5, 4}))
    fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
    fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
    fmt.Println(MaxSequence([]int{-2, 5, 6, -2, -3, 1, 6, -6}))
    fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}