package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

// Factorial is the basic example of Recursive Algorithm
func factorial(num int) int {
	if (num == 1) {
		return 1
	} else {
		return num * factorial(num - 1)
	}
} 