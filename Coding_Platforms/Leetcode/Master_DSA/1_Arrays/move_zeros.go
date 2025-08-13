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
		fmt.Println(moveZeroes(nums))
	}
}

func moveZeroes(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}

	var (
		i, nonZeroIndex int
	)

	// Single loop: move non-zeros to left and swap with zeros
	for i = range nums {
		if nums[i] != 0 {
			// Swap current element with element at nonZeroIndex
			nums[nonZeroIndex], nums[i] = nums[i], nums[nonZeroIndex]
			nonZeroIndex++
		}
	}
	return nums
}
