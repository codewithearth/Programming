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
	word, err = bufio.NewReader(os.Stdin).ReadString('\n') // Take input from user
	if err != nil {
		fmt.Println("Error: unable to read use input:", err)
		os.Exit(1)
	}

	fmt.Println(makeFancyString(word))
}

// func makeFancyString(s string) string {
// 	if len(s) < 3 { // If length of the given string is less then 3 then return.
// 		return s
// 	}

// 	var (
// 		resultWord             strings.Builder
// 		runeWord, rune1, rune2 rune
// 	)

// 	resultWord.Grow(len(s))
// 	for _, runeWord = range s { // Iterate over given string
// 		if runeWord == rune1 && runeWord == rune2 { // If previous both letters are similar to current then continue
// 			continue
// 		}

// 		rune1, rune2 = runeWord, rune1 // Assign current letter to rune1, and previous letter to rune2
// 		resultWord.WriteRune(runeWord) // Arrange the letters
// 	}

// 	return resultWord.String()
// }

func makeFancyString(s string) string {
	if len(s) < 3 { // If length of the given string is less then 3 then return.
		return s
	}

	var (
		resultWord = []byte{s[0]}
		lastWord   = s[0]
		i     int
		count = 1
	)
	for i = 1; i < len(s); i++ { // Iterate over length of string
		if s[i] == lastWord { // Compare last word with current word, and add if count is less then 2
			if count < 2 {
				resultWord = append(resultWord, s[i])
			}
			count++
		} else { // Assign current word to lastWord and count it as 1, add current word into result
			lastWord = s[i]
			count = 1
			resultWord = append(resultWord, s[i])
		}
	}

	return string(resultWord)
}
