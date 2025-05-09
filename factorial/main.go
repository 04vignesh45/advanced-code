package main

import "fmt"

func main() {
	num := 5
	factotial := 1

	for i := 1; i <= num; i++ {
		factotial *= i
	}

	fmt.Println("fac",factotial)
}
 