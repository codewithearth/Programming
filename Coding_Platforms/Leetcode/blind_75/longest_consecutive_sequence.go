package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		err             error
		intVal, res     int
		strData, strVal string
		list            []int
		strList         []string
	)

	strData, err = bufio.NewReader(os.Stdin).ReadString('\n') // Read input from user
	if err != nil {
		log.Println("Error: unable to read data:", err)
		os.Exit(1)
	}

	strList = strings.Fields(strings.TrimSpace(strData)) // Convert input into string list
	list = make([]int, 0, len(strList))

	for _, strVal = range strList { // Convert string list to int list
		if intVal, err = strconv.Atoi(strVal); err == nil {
			list = append(list, intVal)
		}
	}

	res = longestConsecutive(list) // Here we get the Longest Consecutive Sequence

	fmt.Println(res)
}

func longestConsecutive(nums []int) int {
	var (
		currentLen, maxLen, num int
		ok                      bool
		numMap                  = make(map[int]struct{}, len(nums))
	)

	if len(nums) != 0 { // Check if length of list is not 0

		for _, num = range nums { // Convert list into hash value to filter duplicate values
			numMap[num] = struct{}{}
		}

		for num = range numMap {
			if _, ok = numMap[num-1]; !ok { // Start is previous number is not in consecutive sequence
				currentLen = 1

				for { // Iterate over consecutive sequences
					if _, ok = numMap[num+currentLen]; !ok { // If not consecutive sequence then break the loop
						break
					}
					currentLen++
				}

				if maxLen < currentLen { // Compaire to set longest consecutive length
					maxLen = currentLen
				}
			}
		}
	}

	return maxLen
}
