package main

import "fmt"

func main() {

	fmt.Println(equalFrequency("cc"))
}

func equalFrequency(word string) bool {

	list := map[rune]int{}
	list1 := map[int]int{}
	var a, b int
	for _, i := range word {
		list[i]++
	}
	for _, j := range list {
		list1[j]++
	}

	for k, v := range list1 {
		if len(list1) < 2 && (k == 1 || v == 1) {
			return true
		} else if len(list1) <= 2 {
			if a == 0 {
				a, b = k, v
			} else {
				if b == 1 && (a-b == k || a-b == 0) {
					return true
				} else if v == 1 && (k-v == a || k-v == 0) {
					return true
				}
			}
		}
	}
	return false
}
