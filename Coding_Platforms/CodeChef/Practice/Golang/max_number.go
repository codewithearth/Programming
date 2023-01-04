package main

import (
	"fmt"
)

func main() {

	//create variables
	var testcase, size int
	var data string

	fmt.Scan(&testcase) //take input for number of testcases

	for i := 0; i < testcase; i++ { //enter value for each test case

		fmt.Scan(&size)
		fmt.Scan(&data)

		var result int //create variable to store variable

		list := []rune(data) //add test case data into slice

		for idx, val := range list { //iterate over slice

			if idx == 0 { //if value is first

				if val == 49 || len(list) == 1 { // if first number is 1 or size of list is 1 then add 1 in result
					result++
				}
				continue

			} else if list[idx-1] == 48 && val == 49 { //if previous value is 0 and current value is 1 then add 1 to result
				result++

			} else if idx == len(list)-1 && val == 48 { //if last value is 0 then add 1 to result
				result++
			}
		}
		fmt.Println(result)
	}

}
