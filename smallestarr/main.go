package main

import "fmt"

func main() {
	arr := []int{10, 5, 7, 8, 2, 4}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("arrsort", arr)

}
