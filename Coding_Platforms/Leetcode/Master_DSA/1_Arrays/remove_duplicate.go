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
		fmt.Println(removeDuplicates(nums))
	}
}

func removeDuplicates(nums []int) int {

	var (
		i, j int
	)

	//	Two pointer approach
	for i = range nums {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1

}
