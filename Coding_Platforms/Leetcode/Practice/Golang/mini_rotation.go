package main

import "fmt"

func main() {

	fmt.Println(minDominoRotations([]int{1, 2}, []int{2, 1}))
}

func minDominoRotations(tops []int, bottoms []int) int { // Check minimum rotation

	// check the rotation of numbers from both slice to get number
	top_rotation := check_rotation(max_num(tops), tops, bottoms)
	bottom_rotation := check_rotation(max_num(bottoms), bottoms, tops)

	if (top_rotation <= bottom_rotation) && (top_rotation > -1) { // if tops rotation is less or equal to bottom rotation and rotation is greater then -1 then return tops rotation
		return top_rotation

	} else if bottom_rotation > -1 { // if bottoms rotation is greater then -1 then return bottom roatation
		return bottom_rotation

	} else { //else return -1
		return -1
	}

}

func max_num(list []int) int { //check max occurrence of number from slice

	count := map[int]int{}
	var max, val int

	for _, i := range list { //tereate over slice

		if count[i] == 0 { // if new in slice add to map
			count[i] = 1

			if count[i] > max { // if count is greater then max then assign count to max and add that number to val
				max = count[i]
				val = i
			}
		} else { //else increase the count
			count[i] += 1

			if count[i] > max { // if count is greater then max then assign count to max and add that number to val
				max = count[i]
				val = i
			}
		}
	}

	return val
}

func check_rotation(val int, list1 []int, list2 []int) int { //check rotation of numbers

	var count int

	//iterate over list to check rotation
	for idx, i := range list1 {

		if i != val { //if number is not match then rotate from another list
			if list2[idx] == val { //if match from another list then increase count
				count++
			} else { //if not match from another list then return -1
				return -1
			}
		}
	}

	return count
}
