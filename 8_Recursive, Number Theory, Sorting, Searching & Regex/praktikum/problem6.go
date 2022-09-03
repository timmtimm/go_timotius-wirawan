package main

import "fmt"

func merge(a []int, b []int) []int {
    final := []int{}
    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }
    for ; i < len(a); i++ {
        final = append(final, a[i])
    }
    for ; j < len(b); j++ {
        final = append(final, b[j])
    }
    return final
}

func mergeSort(items []int) []int {
    if len(items) < 2 {
        return items
    }
    first := mergeSort(items[:len(items)/2])
    second := mergeSort(items[len(items)/2:])
    return merge(first, second)
}

func MaximumBuyProduct(money int, productPrice []int) {
    countBought := 0
    productPrice = mergeSort(productPrice)
    
    for i:=0; i<len(productPrice); i++ {
        if money < productPrice[i] { break }

        if money >= productPrice[i] {
            money -= productPrice[i]
            countBought++
        }
    }

    fmt.Println(countBought)
}

func main() {
    MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
    MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
    MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
    MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})
    MaximumBuyProduct(0, []int{10000, 30000})
}