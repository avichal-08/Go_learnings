package main

import "fmt"

func main() {
	// fmt.Println("Hello World")

	num := 121

	if isPalindrome(num) {
		fmt.Println(num, "is a palindrome")
	} else {
		fmt.Println(num, "is not a palindrome")
	}
}
