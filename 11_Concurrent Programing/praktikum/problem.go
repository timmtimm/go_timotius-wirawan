package main

import (
	"fmt"
	"bufio"
    "os"
	"strings"
	"sync"
	"runtime"
)

func countCharFromSentence(sentence string, charCount *map[rune]int, wg *sync.WaitGroup, mtx *sync.Mutex)  {
	for _, char := range sentence {
		if char >= 'a' && char <= 'z' {
			(*mtx).Lock()
			var _, isExist = (*charCount)[char]
			
			if isExist {
				(*charCount)[char]++
				} else {
					(*charCount)[char] = 1
			}
			(*mtx).Unlock()
		}
	}

	wg.Done()
}

func main() {
    inputReader := bufio.NewReader(os.Stdin)
	systemThreads := 1
	
	fmt.Printf("Input sentence: ")
    sentence, _ := inputReader.ReadString('\n')
    fmt.Printf("Input systemThreads: ")
	fmt.Scanf("%d", &systemThreads)
	
	lowerSentence := strings.ToLower(sentence)
	sentenceLength := len(lowerSentence)
	countChar := sentenceLength / systemThreads
	lastLetter := sentenceLength % countChar
	
	charCount := map[rune]int{}
	runtime.GOMAXPROCS(systemThreads)
	var wg sync.WaitGroup
	var mtx sync.Mutex
	
	for i:=0; i<systemThreads; i++ {
		var fromChar int = i*countChar
		var untilChar int

		if i == countChar {
			untilChar = i*countChar + lastLetter
		} else {
			untilChar = (i+1)*countChar
		}

		wg.Add(1)
		go countCharFromSentence(lowerSentence[fromChar:untilChar], &charCount, &wg, &mtx)
	}
	
	wg.Wait()

	for key, value := range charCount {
        fmt.Printf("%s: %d\n", string(key), value)
    }
}