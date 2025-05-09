package main

import "fmt"

func main() {
	str := "vignesh"
	sli := []byte(str)
	for i := 0; i < len(sli); i++ {
		for j := 0; j < len(sli)-1; j++ {
			if sli[j] < sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
	fmt.Println("sc", string(sli))
}
