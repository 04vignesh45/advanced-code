package main

import "fmt"

func main() {
	start := 1
	end := 50
	for i := start; i <= end; i++ {
		if primeno(i) {
			fmt.Print(i," ")
		}
	}

}
func primeno(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
