package main

import "fmt"


func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	min, max := largestsmallest(arr)
	fmt.Printf("min %d max %d", min, max)
}


func largestsmallest(arr []int) (min, max int) {
	min, max = arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max

}
