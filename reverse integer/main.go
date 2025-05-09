package main

import "fmt"

func main() {
	arr := "世界"
	ch := []byte(arr)
	for i, j := 0, len(ch)-1; i < j; i, j = i+1, j-1 {
		ch[i], ch[j] = ch[j], ch[i]
	}
	fmt.Println("rev int", string(ch))
}
