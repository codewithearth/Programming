package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		err                    error
		K, T, i                int
		aList                  []string
		aStr, tStr, nkStr, res string
		reader                 *bufio.Reader
	)
	reader = bufio.NewReader(os.Stdin)
	tStr, err = reader.ReadString('\n')
	if err == nil {
		T, err = strconv.Atoi(strings.TrimSpace(tStr))
	}
	if err != nil {
		os.Exit(1)
	}

	for i = 0; i < T; i++ {
		nkStr, err = reader.ReadString('\n')
		if err == nil {
			K, err = strconv.Atoi(strings.Fields(nkStr)[1])
		}
		if err != nil {
			os.Exit(1)
		}

		aStr, err = reader.ReadString('\n')
		if err == nil {
			aList = strings.Fields(aStr)
		}

		res = checkRes(K, aList)
		fmt.Println(res)

	}
}

func checkRes(K int, aBytes []string) string {
	var (
		err  error
		A    int
		aStr string
	)

	sort.Slice(aBytes, func(i, j int) bool {
		a, _ := strconv.Atoi(aBytes[i])
		b, _ := strconv.Atoi(aBytes[j])
		return a < b
	})

	if len(aBytes) == 1 {
		return "YES"
	}else if len(aBytes)==2{

	}

	for _, aStr = range aBytes {
		A, err = strconv.Atoi(aStr)
		if err != nil {
			os.Exit(1)
		}

		if A > K {
			return "NO"
		}
	}

	return "YES"
}
