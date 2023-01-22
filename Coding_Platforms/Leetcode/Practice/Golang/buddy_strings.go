package main

import "fmt"

func main() {

	fmt.Println(buddyStrings("ab", "ba"))
}

func buddyStrings(s string, goal string) bool {

	var num string
	for idx, num := range s {


		if s[idx] != goal[idx] {
			if num == "" {
				num = string(s[idx])
			} else if string(goal[idx]) == num {
				return true
			}
		}else if s[idx] == goal[idx]{


		}
	}
	return false
}
