package main

import "fmt"

func main() {

	fmt.Println(Decode("MCMXC"))
}

func Decode(roman string) int { // decode roman numbers

	//create map for roman number
	roman_data := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	//create two variables for result and to store previous number
	result := 0
	previous_word := ""

	for i := len(roman) - 1; i >= 0; i-- { //iterate over string for reverse it
		for k, v := range roman_data { //iterate over roman data

			if string(roman[i]) == k { //check if string is match with roman data

				if roman_data[previous_word] > v { //check if previous data is bigger-then current data then decrease current data from result
					result -= v

				} else { //if previous data is smaller-then current data then add current data into result
					result += v
				}

				previous_word = k
			}
		}
	}
	return result
}
