package main

import (
	"fmt"
)

func main() {

	fmt.Println(Multiple3And5(10))

}

func Multiple3And5(number int) int { //multiple of 3 and 5

	var result int

	for i := 0; i < number; i++ { // iterate over number

		if i%3 == 0 { //multiples of 3 then add to result
			result += i

		} else if i%5 == 0 { //multiple of 5 then add to result
			result += i
		}
	}

	return result
}
