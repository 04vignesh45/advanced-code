package main

import (
	"fmt"
	"time"
)

// func sender(ch chan<- int) {
// 	for i := 1; i <= 3; i++ {
// 		fmt.Println("Sending:", i)
// 		ch <- i
// 		time.Sleep(time.Second)
// 	}
// 	close(ch)
// }

// func receiver(ch <-chan int) {
// 	for msg := range ch {
// 		fmt.Println("Received:", msg)
// 	}
// }

// func main() {
// 	ch := make(chan int)
// 	go sender(ch)
// 	go receiver(ch)
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Main function exiting.")
// }

func main() {
	ch := make(chan int)

	go sender(ch)
	go receiver(ch)
	time.Sleep(3 * time.Second)
}

func sender(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		fmt.Println("Sen", i)
		ch <- i
		time.Sleep( time.Second)
	}
	close(ch)
}

func receiver(ch <-chan int) {
	for v := range ch {
		fmt.Println("rec", v)
	}
}
