package main

import (
	"fmt"

	"github.com/allaryin/advent-of-code/days"
)

func main() {
	tests := []days.AdventDay{
		&days.Day01{},
		&days.Day02{},
	}

	for n, test := range tests {
		fmt.Printf("--- %d : %s ---\n", n+1, test.Name())
		test.Init()
		if err := test.Run(); err != nil {
			fmt.Printf("!! %v\n", err.Error())
		}
		fmt.Println()
	}
}
