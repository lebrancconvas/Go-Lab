package main

import "fmt"

func main() {
	printName()
}

func printName(firstname string , surname string) {
	fmt.Println(firstname + " " + surname)	
}