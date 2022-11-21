package main

import "fmt"

func main() {

	fmt.Println(FormatDuration(3665))
}

func FormatDuration(seconds int64) string {

	result := ""

	// time := map[int]string{86400: "Day", 3600: "hour", 60: "minute"}

	if seconds >= 86400 {

		fmt.Println(seconds / 86400)
	}

	return result
}
