package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		err     error
		val     int
		word, i string
		nums    []int
		listStr []string
	)
	if word, err = bufio.NewReader(os.Stdin).ReadString('\n'); err == nil {
		listStr = strings.Fields(word)
		for _, i = range listStr {
			if val, err = strconv.Atoi(i); err == nil {
				nums = append(nums, val)
			}
		}
		fmt.Println(majorityElement(nums))
	}
}

func majorityElement(nums []int) int {
	var (
		i, majority int
		elements    = make(map[int]int)
	)

	// count the number of elements in the array
	for i = range nums {
		elements[nums[i]]++
	}

	majority = len(nums) / 2
	// check if the number of elements is greater than the majority
	for i = range elements {
		if elements[i] > majority {
			return i
		}
	}

	return 0
}
