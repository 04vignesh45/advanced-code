package main

import "fmt"

func main() {
	create := map[int]int{
		1: 1,
		2: 2,
		3: 3}
	fmt.Println("map", create)
	var input int
	fmt.Println("enter the number")
	fmt.Scanln(&input)
	check, ok := create[input]
	if !ok {
		fmt.Println("key not present")
	} else {
		fmt.Println("key present", check)
	}
}
