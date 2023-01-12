package main

import "fmt"

func main() {

	fmt.Println(minDominoRotations([]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}))
}

func minDominoRotations(tops []int, bottoms []int) int {

	var top_data, bottom_data int
	tops_val, bottoms_val := max_num(tops, bottoms)

	top_data = check_val(tops_val, tops, bottoms)
	bottom_data = check_val(bottoms_val, bottoms, tops)

	if top_data == bottom_data {
		return top_data
	} else if (top_data < bottom_data) && (top_data > 0) {
		return top_data
	} else if (bottom_data < top_data) && (bottom_data > 0) {
		return bottom_data
	}

	return -1
}

func max_num(tops []int, bottoms []int) (int, int) {

	var max_top, top_val, max_bottom, bottom_val int

	for i := 1; i <= 6; i++ {
		if max := check_count(i, tops); max > max_top {
			max_top = max
			top_val = i
		}
		if max := check_count(i, bottoms); max > max_bottom {
			max_bottom = max
			bottom_val = i
		}
	}
	return top_val, bottom_val
}

func check_count(i int, list []int) int {
	var count int
	for _, j := range list {
		if i == j {
			count++
		}
	}
	return count
}

func check_val(val int, list1 []int, list2 []int) int {

	res_list := []int{}
	var count int
	for idx, i := range list1 {
		if i != val {
			if list2[idx] == val {
				res_list = append(res_list, list2[idx])
				count++
			} else {
				return -1
			}
		} else {
			res_list = append(res_list, i)
		}
	}

	return count

}
