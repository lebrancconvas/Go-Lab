package main

import "fmt"

var a int
var b int

func main(){
	fmt.Println(addNumber(1,2))
	fmt.Println(addNumber(1,8))
	fmt.Println(addNumber(9,10))
}

func addNumber(a int,b int) int {
	return a + b
}