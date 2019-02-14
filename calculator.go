package main

import "./math"
import "fmt"

func main() {
	plus := math.Plus(1,2)
	div := math.Divide(4,2)
	fmt.Println(div,plus)
}