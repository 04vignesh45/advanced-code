package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	sli := []int{1, 2, 3, 4, 5}
	array(&arr)
	slice(sli)
	fmt.Println("Arr", arr)
	fmt.Println("sli", sli)

}

func array(arr *[5]int) {
	for index, value := range arr {
		arr[index] = value * 2
	}
	fmt.Println("resul", arr)
}

func slice(arr []int) {
	for index, value := range arr {
		arr[index] = value * 2
	}
	fmt.Println("resul", arr)
}
