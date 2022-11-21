package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(FormatDuration(1))
}

func FormatDuration(seconds int64) string { //convert seconds to accurate time

	//create slice for store years, days, hours, minutes and seconds
	//create variable to store result in string format
	data := []string{}
	result := ""

	if seconds >= 31536000 { //if seconds in greater then 31536000 then add it as years

		//check how many yearss we got
		//and update rest of seconds
		year := int(seconds / 31536000)
		seconds = seconds % 31536000

		if year > 1 { //if year is greater then one then write as years
			data = append(data, strconv.Itoa(year)+" years")
		} else { //else write as year
			data = append(data, strconv.Itoa(year)+" year")
		}
	}

	if seconds >= 86400 { //if seconds in greater then 86400 then add it as days

		//check how many days we got
		//and update rest of seconds
		day := int(seconds / 86400)
		seconds = seconds % 86400

		if day > 1 { //if day is greater then one then write as days
			data = append(data, strconv.Itoa(day)+" days")
		} else { //else write as day
			data = append(data, strconv.Itoa(day)+" day")
		}
	}

	if seconds >= 3600 { //if seconds in greater then 3600 then add it as hours

		//check how many hours we got
		//and update rest of seconds
		hour := int(seconds / 3600)
		seconds = seconds % 3600

		if hour > 1 { //if hour is greater then one then write as hours
			data = append(data, strconv.Itoa(hour)+" hours")
		} else { //else write as hour
			data = append(data, strconv.Itoa(hour)+" hour")
		}
	}

	if seconds >= 60 { //if seconds in greater then 60 then add it as minute

		//check how many minutes we got
		//and update rest of seconds
		minute := int(seconds / 60)
		seconds = seconds % 60

		if minute > 1 { //if minutes is greater then one then write as minutes
			data = append(data, strconv.Itoa(minute)+" minutes")
		} else { //else write as minute
			data = append(data, strconv.Itoa(minute)+" minute")
		}
	}

	if seconds < 60 && seconds > 0 { //if seconds in greater then 0 and less then 60 then add it as minutes

		if seconds > 1 { //if seconds is greater then one then write as seconds
			data = append(data, strconv.Itoa(int(seconds))+" seconds")
		} else { //else write as second
			data = append(data, strconv.Itoa(int(seconds))+" second")
		}
	}

	if len(data) > 1 { //if lenght of data is greater then one then write data in proper format and add to result

		before_result := strings.Join(data[:len(data)-1], ", ")
		result = before_result + " and " + data[len(data)-1]
	} else { //else add direct data to result
		result = strings.Join(data, "")
	}
	return result
}
