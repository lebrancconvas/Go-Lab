package main

import "fmt"

func main() {
	fmt.Print("Total Summation is ")
	sumAllnum(1,2,3,4,5,6,7,8,9,10)
}

func sumAllnum(num...int) {
	var total int
	for _ , n := range num {
		total += n
	}
	fmt.Println(total)
} 