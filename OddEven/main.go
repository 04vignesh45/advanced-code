package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	odd := make(chan int)
	even := make(chan int)
	wg.Add(2)
	go oddch(odd, even, &wg)
	go evench(odd, even, &wg)
	wg.Wait()

}

func oddch(odd, even chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		odd <- i
		fmt.Println("odd", i)
		<-even
	}
}

func evench(odd, even chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-odd
		fmt.Println("even", i)
		even <- i
	}
}
