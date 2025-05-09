package main

import "fmt"

func main() {
	arr := []int{1, 23, 4, 7, 0, 5}
	for i:=0 ;i<len(arr); i++{
		for j:=0 ;j<len(arr)-1; j++{
			if arr[j]>arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("arr",arr)
}