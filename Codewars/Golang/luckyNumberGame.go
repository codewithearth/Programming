package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//Create function for lucky number
func luckyNumber(total_numbers, alice, bob int) {
	a, b := 0, 0
	for j := 0; j < total_numbers; j++ {
		wg.Add(1)
		go func(alice, bob, a, b int) {
			defer wg.Done()
			var value int
			fmt.Scan(&value)
			if value == 0 {
				return
			} else if value%alice == 0 {
				a += 1
			} else if value%bob == 0 {
				b += 1
			}

		}(alice, bob, a, b)
	}
	wg.Wait()

	if a > b || a == b {
		fmt.Println("BOB")
	} else {
		fmt.Println("ALICE")
	}
}

func main() {
	var total_turns, total_numbers, alice, bob int
	fmt.Scan(&total_turns)
	for i := 0; i < total_turns; i++ {
		fmt.Scan(&total_numbers, &alice, &bob)
		luckyNumber(total_numbers, alice, bob)
	}
}



