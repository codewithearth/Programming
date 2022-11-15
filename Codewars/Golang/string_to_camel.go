package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
}

func ToCamelCase(s string) string { // convert normal string to camel case

	// create two variables, one for result and one for flag
	result := ""
	upper := 0

	for _, i := range s { //iterate over string

		if i == 45 || i == 95 { //if string is underscore or hypen then continue for next iteration and add 1 into flag
			upper = 1
			continue

		} else if upper == 1 { //if flag is 1 then uppercase the string and continue for nexr iteration and add 0 into flag

			result += strings.ToUpper(string(i))
			upper = 0
			continue
		}

		result += string(i)

	}

	return result
}
