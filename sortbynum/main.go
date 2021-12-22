package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "is2 Thi1s T4est 3a" //This is a Test

	fmt.Println(sortbynum(str))
}

func sortbynum(str string) string {
	tstr := strings.Split(str, " ")
	nstr := ""
	var nums = []string{"1", "2", "3", "4"}

	for j := 0; j < len(nums); j++ {
		for _, i := range tstr {
			if strings.Contains(i, nums[j]) {
				nstr += strings.Replace(i, nums[j], "", -1) //Removing numbers from string and add string to new
				nstr += " "
			}
		}
	}

	return nstr
}
