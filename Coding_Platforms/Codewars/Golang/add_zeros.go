package main

import "fmt"

func main() {

	fmt.Println(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}))
}

func MoveZeros(arr []int) []int { //Move all zeros to the end

	//create two slice,
	//one is for result silce and another is for store all zeros
	result := []int{}
	zeros := []int{}

	for _, i := range arr { //iterate over the slice

		if i == 0 { //if if interger is zero then add it to zeros slice
			zeros = append(zeros, 0)
			continue
		}
		//append add non-zero integer to the result
		result = append(result, i)
	}

	//finally we have result slice of non-zeros
	//then add zeros slice in the end of result
	result = append(result, zeros...)

	return result
}
