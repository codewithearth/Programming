package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Println(SumOfDivided([]int{-15, -30, -45, -65, -85}))

}

func SumOfDivided(lst []int) string { //sum of divided

	//create two variables
	result := ""
	max_value := Max(lst)
	for i := 2; i < max_value; i++ { //iterate for prime number

		if IsPrime(i) { //check prime number

			if a, b := CheckValue(i, lst); b == false { //check value is divided by prime number,if not divided by prime number then continue
				continue

			} else { //if divided then add prime number and sum of values to result
				result += "(" + strconv.Itoa(i) + " " + strconv.Itoa(a) + ")"
			}
		}
	}

	return result
}

func Max(lst []int) int { //check maximum number form the list

	num := 0
	for _, i := range lst {
		j := math.Abs(float64(i))
		if num < int(j) {
			num = int(j)
		}
	}
	return num
}

func IsPrime(num int) bool { //check prime number

	for i := num - 1; i > 1; i-- {

		if num%i == 0 { //if number is divided by any number then return false
			return false
		}
	}
	return true
}

func CheckValue(i int, lst []int) (int, bool) { //check if value divided by prime number then return total of value

	result := 0
	is_divide := false
	for _, j := range lst {

		if j%i == 0 {
			result += j
			is_divide = true
		}
	}

	return result, is_divide
}
