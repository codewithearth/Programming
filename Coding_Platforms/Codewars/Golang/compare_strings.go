package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println(InArray(
		[]string{"live", "arp", "strong"},
		[]string{"lively", "alive", "harp", "sharp", "armstrong"},
	))

}

func InArray(array1 []string, array2 []string) []string { //compare strings from two arrays

	//create a variable to store result
	result := []string{}

	for _, data1 := range array1 { //iterate over first array
		for _, data2 := range array2 { //iterate over second array

			if strings.Contains(data2, data1) && !Contains(result, data1) { //if string of 1st array is a subset of string of 2nd array,

				result = append(result, data1)
				break
			}
		}
	}

	//sort the result to get lexicographical order result
	sort.Strings(result)
	return result

}

func Contains(s []string, str string) bool { //check duplicate value

	for _, v := range s { //iterate over slice
		if v == str { //if slice already contain string
			return true
		}
	}

	return false
}
