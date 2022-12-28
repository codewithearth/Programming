package main

import (
	"fmt"
	"strconv"
)

func main() {

	var testcase, size int
	var data string
	var value = []int{}
	fmt.Scan(&testcase)
	for i := 0; i < testcase; i++ {
		fmt.Scan(&size)
		fmt.Scan(&data)
		for _, j := range data {
			val, _ := strconv.Atoi(string(j))
			value = append(value, val)

		}
		fmt.Println(CheckMax(value))
	}

}

func CheckMax(list []int) int {

	result := 1

	for _, k := range list {
		fmt.Println(k)
	}
	return result
}
