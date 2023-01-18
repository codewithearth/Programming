package main

import "fmt"

func main() {

	fmt.Println(canBeIncreasing([]int{105, 924, 32, 968}))
}

func canBeIncreasing(nums []int) bool { //check increasing number

	var nums1 = make([]int, len(nums))

	if check_increase(nums) { // if increasing then return true
		return true

	} else {
		for i, _ := range nums { // range over slice 

			//copy slice into another slice and remove one number every time and check increase order
			if copy(nums1, nums);check_increase(append(nums1[:i], nums1[i+1:]...)) {
				return true
			}
		}
	}
	return false
}

func check_increase(num []int) bool { // check the increase order

	for idx, i := range num {//iterate over slice

		// if previous number is greater then current number then return false
		if (idx > 0) && (num[idx-1] >= i) { 
			return false
		}
	}
	return true
}
