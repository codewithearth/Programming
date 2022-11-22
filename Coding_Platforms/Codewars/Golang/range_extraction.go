package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(Solution([]int{1, 2, 3, 6, 7, 8, 9, 10, 11, 20, 22, 23, 31, 33}))
}

func Solution(list []int) string { //range extraction from list

	//create three variables
	//one for result ,one for last_value and one as flag
	result := ""
	var last_value, count int

	for idx, data := range list { //iterate over list

		if idx >= 1 { //if index is greater then one
			if (last_value+1) == data || (last_value-1) == data { //if data of list match with +1 or -1 of previous data then add current data into last_value and count increases by +1 and continue the iteration

				last_value = data
				count += 1
				continue
			}

		}

		if count == 1 { //if count is 1 then add "," in result

			result += ","

		} else if count == 2 { //if count is 2 then add ",", previous data and "," in result

			result += "," + strconv.Itoa(last_value) + ","

		} else if count > 2 { //if count is greater then 2 then add "-", previous data and "," in result

			result += "-" + strconv.Itoa(last_value) + ","
		}

		//assign 1 to count and current data to last_data
		count = 1
		last_value = data

		//add last data into result
		result += strconv.Itoa(last_value)

	}

	// if count is 2 or greater then 2 then add "-" or "," and last value to result because it miss in last value when data in range
	if count == 2 {
		result += "," + strconv.Itoa(last_value)

	} else if count > 2 {
		result += "-" + strconv.Itoa(last_value)
	}

	return result
}
