package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(duplicate_count("ABBA"))
}

func duplicate_count(s1 string) int { //count duplicate

	//create two variables, one is for result and another is for store all letters in lower case
	result := 0
	str1 := strings.ToLower(s1)

	for _, i := range strings.ToLower(s1) { //iterate over string

		if strings.Count(str1, string(i)) >= 2 { //if letter count is more then 1 then add it to duplicate and remove that letter from string
			result += 1
			str1 = strings.Replace(str1, string(i), "", -1)
		}
	}

	return result
}
