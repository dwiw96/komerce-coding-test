package psbb

import (
	"fmt"
)

func MinimumBus(n int, members []int) {
	passengerMap := make(map[int]int)

	var bus int

	for _, v := range members {
		passengerMap[v]++
	}

	for _, val := range members {
		target := 4 - val
		if passengerMap[val] <= 0 {
			continue
		}
		if passengerMap[target] > 0 {
			passengerMap[val]--
			passengerMap[target]--
		} else {
			passengerMap[val]--
		}
		bus++
	}

	fmt.Printf(">> Minimum bus required is: %d", bus)
}
