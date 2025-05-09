package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	// arr2 := make([]int, len(arr))
	arr2:=arr
	arr2[0]=11
	fmt.Println(arr)

	// copy(arr2, arr)
	// arr2[0]=10000
	// fmt.Println("Source:", arr)
	// fmt.Println("Destination after copy:", arr2)
	// fmt.Println("Destination after copy:", arr)

}
