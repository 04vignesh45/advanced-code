package main

import "fmt"

func main() {

	var a int
	fmt.Println("enter the number")
	fmt.Scan(&a)
	pass := fibo(a)
	fmt.Println("fibo series", pass)
}

func fibo(a int) []int {
	create := make([]int, a)
	create[0], create[1] = 0, 1
	for i := 2; i < a; i++ {
		create[i] = create[i-1] + create[i-2]
	}
	return create
}
