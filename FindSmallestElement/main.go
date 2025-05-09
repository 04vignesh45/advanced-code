package main

import (
	"fmt"
	"sort"
)

// var third int = 3

func main() {
	var size int
	fmt.Println("enter the numbers")
	fmt.Scanln(&size)
	fmt.Println("your array elements")
	create := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&create[i])
	}
	fmt.Println("array elements are", create)
	pass := thirdsmallest(create)
	fmt.Println("third smallest element:=>", pass)
}
func thirdsmallest(a []int) int {
	sort.Ints(a)
	return a[2-1]
}

