package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(SpinWords("i love pizza and burger"))
}

func SpinWords(str string) string { //spinning words

	//create two
	word_list := strings.Split(str, " ")
	result_list := []string{}

	for _, word := range word_list {

		if len(word) > 4 {

			reverse_word := ""
			for i := len(word); i > 0; i-- {

				reverse_word += string(word[i-1])
			}
			fmt.Println(reverse_word)

			result_list = append(result_list, reverse_word)
		} else {
			result_list = append(result_list, word)
		}

	}

	return strings.Join(result_list, " ")
}
