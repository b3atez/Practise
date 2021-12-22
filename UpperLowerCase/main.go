package main

import (
	"fmt"
)

func main() {
	var (
		str string = "HeLlo JOhn! hOw ARe yoU" // 8 Upper  10 Lower cases
	)

	ucase, lcase := UpperLowerCount(str)

	fmt.Printf("Upper cases: %d\n", ucase)
	fmt.Printf("Lower cases: %d\n", lcase)

}

func UpperLowerCount(str string) (int, int) { //This function count only Upper or Lower cases not symbols
	var ucase, lcase int

	for _, val := range str {
		if val >= 65 && val <= 90 {
			ucase++
		}
		if val >= 97 && val <= 122 {
			lcase++
		}
	}

	return ucase, lcase

}
