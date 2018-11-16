package main

import (
	"fmt"

	"github.com/mike-dunton/flotto/lucky-money"
)

func main() {
	for _, winner := range lucky.Results([5]int{1, 2, 3, 4, 5}) {
		fmt.Printf("Date: %s,  Numbers: %d , Winners: %s, Prize: %s\n", winner.Date, winner.Numbers, winner.Winners, winner.Prize)
	}
}
