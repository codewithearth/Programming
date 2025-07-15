package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		err  error
		word string
	)

	word, err = bufio.NewReader(os.Stdin).ReadString('\n') // Get input from user
	if err != nil {
		fmt.Println("Error: unable to get input:", err)
		os.Exit(1)
	}

	fmt.Println(isValid(word))
}

func isValid(word string) bool {
	var (
		asciiVal              int
		i                     rune
		isVowels, isConsonant bool
		vowels                = map[string]bool{"A": true, "E": true, "I": true, "O": true, "U": true, "a": true, "e": true, "i": true, "o": true, "u": true}
	)

	for _, i = range word { // Iterate over the word to match all conditions

		asciiVal = int(i)
		if vowels[string(i)] { // Check for Vowels
			isVowels = true
		} else if (asciiVal > 64 && asciiVal < 91) || (asciiVal > 96 && asciiVal < 123) { // Check for Consonant
			isConsonant = true
		} else if asciiVal < 48 || asciiVal > 57 { // Check for digits, If not then return false
			return false
		}
	}

	if isVowels && isConsonant && len(word) > 2 { // Check all criteria for validate word
		return true
	}

	return false
}
