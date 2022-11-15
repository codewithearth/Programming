package main

import "fmt"

func main() {

	fmt.Println(Solution("abc"))
}

func Solution(str string) []string { //split string

	var result []string
	var data string

	if len(str)%2 != 0 { //add underscore if len of string is not even
		str = str + "_"
	}

	for idx, i := range str { //iterate over string

		data = data + string(i)

		if idx%2 != 0 { //add data into slice

			result = append(result, data)
			data = ""
		}
	}
	return result
}
