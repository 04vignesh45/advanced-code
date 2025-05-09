package main

import (
	"fmt"
	"sync"
	"unicode"
)

func main() {
	separateCase := "GoLangProgram"
	var wg sync.WaitGroup
	find := make(chan rune)
	check := make(chan bool)
	wg.Add(2)
	go SmallCase(find, check, &wg)
	go CapsCase(find, check, &wg)
	for _, char := range separateCase {
		find <- char
		<-check
	}
	close(find)
	close(check)
	wg.Wait()
}

func SmallCase(find chan rune, check chan bool, wg *sync.WaitGroup) {
	for findcase := range find {
		select {
		case <-check:
			find <- findcase
		default:
			if unicode.IsLower(findcase) {
				fmt.Println("small case", string(findcase))
				check <- true
			} else {
				find <- findcase
			}
		}
	}
	defer wg.Done()
}

func CapsCase(find chan rune, check chan bool, wg *sync.WaitGroup) {
	for findcase := range find {
		select {
		case <-check:
			find <- findcase
		default:
			if unicode.IsUpper(findcase) {
				fmt.Println("upper case", string(findcase))
				check <- true
			} else {
				find <- findcase
			}
		}
	}
	defer wg.Done()

}
