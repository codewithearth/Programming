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
		err     error
		val     int
		word, i string
		nums    []int
		listStr []string
	)
	if word, err = bufio.NewReader(os.Stdin).ReadString('\n'); err == nil {
		listStr = strings.Fields(word)
		for _, i = range listStr {
			if val, err = strconv.Atoi(i); err == nil {
				nums = append(nums, val)
			}
		}
		fmt.Println(maxProfit(nums))
	}
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var (
		price, minPrice, maxProfit int
	)

	minPrice = prices[0]

	// Single pass: track minimum price and calculate potential profit
	for _, price = range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}
