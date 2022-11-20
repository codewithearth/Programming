package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(SpinWords("i love pizza and burger"))
}

func SpinWords(str string) string { //spinning words

	//create two slices,
	//one for store words as list and another for store result
	word_list := strings.Split(str, " ")
	result_list := []string{}

	for _, word := range word_list { //iterate over word list

		if len(word) > 4 { //if word length is more then four then reverse the word

			//create a veriable to store reverse variable
			reverse_word := ""

			for i := len(word); i > 0; i-- { //iterate over length of word and store reverse letters into reverse_word

				reverse_word += string(word[i-1])
			}

			//append reverse word into result_list
			result_list = append(result_list, reverse_word)

		} else { // if word length is less then five then simply add word to result_list

			result_list = append(result_list, word)
		}

	}

	//join words from slice with one space
	return strings.Join(result_list, " ")
}
