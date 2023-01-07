package main

import "fmt"

func main() {

	//create a veriable for test cases and take input
	var test_case int
	fmt.Scanln(&test_case)

	for i := 0; i < test_case; i++ { //iterate for testcases and print result

		fmt.Println(AvgNumber())
	}
}

func AvgNumber() int { //check average number

	//create variable for n, k, v, data and result and take input for n, k, and v.
	var n, k, v, data, result int
	fmt.Scanln(&n, &k, &v)

	for j := 0; j < n; j++ { //iterate for n times to take value of n.

		//create variable and take value for variable and add that value in data
		var val int
		fmt.Scan(&val)
		data += val
	}

	V := (n + k) * v //store sum of all data into V.

	if ((V - data) % k) == 0 { //if result is not integer

		if result = (V - data) / k; result > 0 { //store average number into result if average is greater then 0 then return result
			return result
		}
	}
	return -1
}
