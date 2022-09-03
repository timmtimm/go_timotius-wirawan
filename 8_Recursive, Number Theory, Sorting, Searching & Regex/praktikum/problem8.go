package main

import "fmt"

type pair struct {
    name string
    count int
}

func merge(a []pair, b []pair) []pair {
    final := []pair{}
    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if a[i].count > b[j].count {
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

func mergeSort(items []pair) []pair {
    if len(items) < 2 {
        return items
    }
    first := mergeSort(items[:len(items)/2])
    second := mergeSort(items[len(items)/2:])
    return merge(first, second)
}

func MostAppearItem(items []string) []pair {
    counts := []pair{}

    for _, item := range items {
        itemFound := false
        
        for i:=0; i<len(counts); i++ {
            if item == counts[i].name {
                itemFound = true
                counts[i].count++
                break
            }
        }

        if itemFound == false {
            counts = append(counts, pair{name: item, count: 1})
        }
    }

    return mergeSort(counts)
}


func main() {
    fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
    fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
    fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}