package main

import (
	"fmt"
)

func main() {
	var total_turns, total_numbers, bob, alice int
	fmt.Scan(&total_turns)
	for i := 0; i < total_turns; i++ {
		fmt.Scan(&total_numbers, &bob, &alice)
		fmt.Println(LuckyNumber(total_numbers, bob, alice))
	}
}

func LuckyNumber(total_numbers, bob, alice int) string { // check lucky number

	var value int
	var data = []int{}

	a := 0
	for j := 0; j < total_numbers; j++ {
		fmt.Scanf("%d", &value)
		if value == 0 {
			continue
		} else {
			data = append(data, value)
		}
	}
	for {
		if a == 0 {
			for idx_i, i := range data {

				if i%bob == 0 {
					a = 1
					data = append(data[:idx_i], data[idx_i+1:]...)
				}
			}
			if a != 1 {
				return "ALICE"
			}
		} else if a == 1 {
			for idx_j, j := range data {

				if j%alice == 0 {
					a = 0
					data = append(data[:idx_j], data[idx_j+1:]...)
				}
			}
			if a != 0 {
				return "BOB"
			}
		}
	}
}
