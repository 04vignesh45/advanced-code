package main

import "fmt"

func main() {
	a := 120
	if palindrom(a) {
		fmt.Println(a, "is a palindrome")
	} else {
		fmt.Println(a, "is not a palindrome")

	}
}

func palindrom(n int) bool {
	rev, temp := 0, n
	for temp > 0 {
		rev = rev*10 + temp%10
		temp /= 10
	}
	return n == rev
}
