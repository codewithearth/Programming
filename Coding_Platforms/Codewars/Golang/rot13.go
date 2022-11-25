package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(Rot13("EBG13 rknzcyr."))
}

func Rot13(msg string) string { //convert rot values

	//create two variables
	//one for result and one is map that store rot value
	result := ""
	rot_value := map[string]string{"a": "n", "b": "o", "c": "p", "d": "q", "e": "r", "f": "s", "g": "t", "h": "u", "i": "v", "j": "w", "k": "x", "l": "y", "m": "z"}

	for _, letter := range msg { //iterate over msg

		if string(letter) == strings.ToLower(string(letter)) { //if letter is lowercase

			if word := CheckValue(string(letter), rot_value); word != "" { //if letter is in rot_value then return its rot_value to word and add to result
				result += word

			} else { // if letter is in not in rot_value then add that letter same as it in result
				result += string(letter)
			}

		} else { //if word is in uppercase

			// then convert uppercase letter to lower case first to find its rot_value and add its rot_value to result in uppercase
			letter := strings.ToLower(string(letter))
			result += strings.ToUpper(string(CheckValue(string(letter), rot_value)))
		}
	}
	return result
}

func CheckValue(letter string, rot_value map[string]string) string { //check letter in rot_value

	for i, j := range rot_value { //iterate over rot_value

		if letter == i { //if letter is equal to key of rot_value then return value of that key
			return j
		} else if letter == j { //if letter is equal to value of rot_value then return key of that value
			return i
		}

	}

	//otherwise return empty string
	return ""
}
