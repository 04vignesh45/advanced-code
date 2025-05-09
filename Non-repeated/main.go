package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6,10}
	nonrep := findnonrep(arr)
	fmt.Println("Non-repeating elements:", nonrep)
}

func findnonrep(arr []int) ([]int, bool) {
	create := make(map[int]int)
	for _, v := range arr {
		create[v]++
	}
	var nonrep []int
	for _, v := range arr {
		if create[v] == 1 {
			nonrep = append(nonrep, v)
		}
	}
	return nonrep, len(nonrep) > 0
}
