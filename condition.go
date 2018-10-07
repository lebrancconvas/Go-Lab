package main

import "fmt"

func main() {
	var prisonnumber int
	fmt.Print("How Many Prison in this jail : ")
	fmt.Scanf("%d",&prisonnumber)
	if prisonnumber > 5 {
		fmt.Println("OK!")
	} else {
		fmt.Println("Someone lost!!, Find him!")
		fmt.Println("SHIT!")
	}
}