package main

import "fmt"

func main() {
	rate := 5
	switch rate {
		case 1 : fmt.Println("Case not good")
		case 2 : fmt.Println("Case is good")
		default : fmt.Println("Case Validate")
	}
}