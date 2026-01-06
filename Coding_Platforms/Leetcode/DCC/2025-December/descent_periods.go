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
		err        error
		num        int
		input, val string
		prices     []int
	)

	input, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		os.Exit(1)
	}

	for _, val = range strings.Fields(input) {
		num, _ = strconv.Atoi(val)
		prices = append(prices, num)
	}

	fmt.Println(getDescentPeriods(prices))
}

func getDescentPeriods(prices []int) int64 {
	var (
		i         int
		result, n = 1, 1
	)

	for i = 1; i < len(prices); i++ {
		if prices[i-1]-prices[i] == 1 {
			n++
		} else {
			n = 1
		}
		result += n
	}

	return int64(result)
}