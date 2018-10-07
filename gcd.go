package main

import "fmt"

func main() {
	add := addNumber(10,10)
	remove := removeNumber(10,10)
	multi := multiNumber(10,10)
	divide := divideNumber(10,10) 
	fmt.println(add,remove,multi,divide)
}

func addNumber(a int , b int) {
	return a + b
}

func removeNumber(a int , b int) {
	return a - b
}

func multiNumber(a int , b int) {
	return a * b
}

func divideNumber(a int , b int) {
	return a / b
}