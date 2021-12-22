package main

import (
	"fmt"
)

func main() {
	var str string = "Hello! How are You! Are you ok?"

	point := countery(str)

	for i, val := range point {
		fmt.Printf("%d - o %d index\n", i+1, val)
	}
}

func countery(str string) []int { // find how many 'o' and its index
	var (
		point []int
	)
	for i, val := range str {
		if val == 111 {
			point = append(point, i)
		}
	}

	return point
}
