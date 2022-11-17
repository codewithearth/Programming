package main

import "fmt"

func main() {

	fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
}

func LongestConsec(strarr []string, k int) string { //check longest consecutive string from slice

	//create two variables, one for store result and another is for check maximum length
	result := ""
	max_len := 0

	if len(strarr) < k { //if length of strings in slice is smaller then consecutive length

		return result
	} else {

		for i := 0; i < len(strarr)-(k-1); i++ { //iterate over the length of slice

			//create a variable to store consecutive word
			word := ""

			for j := i; j < k+i; j++ { //iterate over consecutive length and store result in word

				word += strarr[j]
			}

			//if length of word is bigger then all previous word then,
			//store that word in result and assign length of that word to max_len
			if len(word) > max_len {

				result = word
				max_len = len(word)
			}
		}
	}
	return result
}
