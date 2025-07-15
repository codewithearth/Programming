package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		T, zeros, ones  int
		S, nStr, res, tStr string
		reader             *bufio.Reader
	)
	reader = bufio.NewReader(os.Stdin)

	tStr, _ = reader.ReadString('\n')
	T, _ = strconv.Atoi(strings.TrimSpace(tStr))
	for range T {
		var rCount int
		reader.ReadString('\n')
		strconv.Atoi(strings.TrimSpace(nStr))

		S, _ = reader.ReadString('\n')
		zeros = strings.Count(S, "0")
		ones = strings.Count(S, "1")

		if zeros >= ones && ones != 0 {
			rCount = ones
		} else if zeros != 0 && ones != 0 {
			rCount = zeros
		}
		res = checkWinner(rCount)
		fmt.Println(res)
	}
}

func checkWinner(rCount int) string {
	if rCount%2 == 0 {
		return "Ramos"
	} else {
		return "Zlatan"
	}
}
