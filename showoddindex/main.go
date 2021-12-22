package main

import (
	"fmt"
)

func main() {
	var (
		str string
	)

	fmt.Scan(&str)

	for i, val := range str { //Get odd index from string
		if i%2 != 0 {
			fmt.Print(string(val))
		}
	}
}
