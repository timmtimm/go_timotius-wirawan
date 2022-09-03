package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
    max, indexMax := 0, 0
    
    for index, card := range cards {
        possible := card[0] == deck[0] || card[0] == deck[1] || card[1] == deck[0] || card[1] == deck[1]
        sum := card[0] + card[1]

        if possible && sum > max {
            max = sum
            indexMax = index
        }
    }

    if max == 0 && indexMax == 0 {
        return "tutup kartu"
    }

    return cards[indexMax]
}

func main() {
    fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
    fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
    fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}