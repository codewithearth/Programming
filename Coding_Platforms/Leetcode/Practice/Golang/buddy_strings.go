package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		err     error
		s, goal string
		reader  *bufio.Reader
	)

	reader = bufio.NewReader(os.Stdin) // Create new reader to read from user
	s, err = reader.ReadString('\n')   // Read input of s
	if err != nil {
		fmt.Println("Error: unable to read input:", err)
		os.Exit(1)
	}

	goal, err = reader.ReadString('\n') // Read input of goal
	if err != nil {
		fmt.Println("Error: unable to read input:", err)
		os.Exit(1)
	}

	fmt.Println(buddyStrings(strings.TrimSpace(s), strings.TrimSpace(goal)))

}

func buddyStrings(s string, goal string) bool {
	var (
		idx, preIdx, wordLen int
		res, ok              bool
		sRune, swapWord      rune
		runes                []rune
		wordMap              = make(map[rune]int)
	)

	runes = []rune(s)
	wordLen = len(s)
	if wordLen > 1 && (wordLen == len(goal)) { // Check the lenght of both words
		for idx, sRune = range s { // Iterate over one word

			if !res {
				if _, ok = wordMap[sRune]; !ok { // Check duplicate letters
					wordMap[sRune]++
				} else {
					res = true
				}
			}

			if sRune != rune(goal[idx]) { // Check mismatch letters
				if swapWord == rune(0) {
					swapWord = sRune
					preIdx = idx
					continue
				} else if swapWord == rune(goal[idx]) { // Swap letters and compare the results
					runes[preIdx], runes[idx] = sRune, runes[preIdx]
					if string(runes) == goal {
						return true
					}
				}
				return false
			}
		}
		if swapWord != rune(0) { // Check is only one word is mismatched
			return false
		}
	}
	return res
}
