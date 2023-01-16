package main

import "fmt"

func main() {

	fmt.Println(equalFrequency("abcde"))
}

func equalFrequency(word string) bool {// make equal frequency 

	list := map[rune]int{}
	list1 := map[int]int{}
	var a, b int

	for _, i := range word {// store occurrence of word in map
		list[i]++
	}

	for _, j := range list {// store occurrence of count in map
		list1[j]++
	}

	for k, v := range list1 { //iterate over list1

		if len(list1) < 2 && (k == 1 || v == 1) { // if length is less then 2 and one of the value is 1 then return true
			return true

		} else if len(list1) == 2 { // if length is 2

			if a == 0 { // if a is empty then store k, v in a and b
				a, b = k, v

			} else {
				if b == 1 && (a-b == k || a-b == 0) { //if b is 1 and subtraction of a,b is k or 0 then return true 
					return true

				} else if v == 1 && (k-v == a || k-v == 0) {//if v is 1 and subtraction of k,v is a or 0 then return true
					return true
				}
			}
		}
	}
	return false
}
