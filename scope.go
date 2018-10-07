package main

import "fmt"

var gVariable string = "Global State"

func main() {
	lVariable := "Local State"
	fmt.Println("In Main Status :",gVariable)
	fmt.Println("In Main Status :",lVariable)
	printScope()
}

func printScope() {
	lfVariable := "Local Function State"
	fmt.Println("In Function State:",gVariable)
	// fmt.Println("In Function State:",lVariable)
	fmt.Println("In Function State Local:",lfVariable)
	fmt.Println("In Function State Local:",lfVariable)
}

