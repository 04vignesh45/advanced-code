// package main

// import "fmt"

// func main() {
// 	fmt.Println("enter array size")
// 	var size int
// 	fmt.Scanln(&size)
// 	fmt.Println("enter your array elements")
// 	create := make([]int, size)
// 	for i := 0; i < size; i++ {
// 		fmt.Scanln(&create[i])
// 	}
// 	fmt.Println("Array elements are", create)

// }

package main

import "fmt"

func main() {

	fmt.Println("enter array size")
	var size int
	fmt.Scanln(&size)
	fmt.Println("enter array elements")
	create := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&create[i])
	}
	fmt.Println(create)
}
