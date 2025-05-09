package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	sum := 0

	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
}
