package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		err   error
		input string
	)

	input, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(countPalindromicSubsequence(input))
}

func countPalindromicSubsequence(s string) int {
	var (
		letter        rune
		uniqueLetter  string
		result        = make(map[string]struct{})
		uniqueLetters = make(map[string]struct{})
	)

	for _, letter = range s {
		uniqueLetters[string(letter)] = struct{}{}
	}

	for uniqueLetter = range uniqueLetters {
		first := strings.Index(s, uniqueLetter)
		last := strings.LastIndex(s, uniqueLetter)

		if first < last {
			var midLetter string
			uniqueMidLetters := make(map[string]struct{})
			for i := first + 1; i < last; i++ {
				uniqueMidLetters[string(s[i])] = struct{}{}
			}

			for midLetter = range uniqueMidLetters {
				result[uniqueLetter+midLetter+uniqueLetter] = struct{}{}
			}
		}

	}

	return len(result)
}
