package main

import (
	"fmt"
	"math/rand"
	"slices"
)

// 3x3 board
// type board struct {
// 	line [15]string
// }

func main() {
	// basic 3x3 board
	for range 2 {
		fmt.Println(" ")
	}

	positions := []int{2, 8, 13}

	xo := []string{"X", "O"}

	hLine(positions, xo)
	xLine(positions, xo)
	hLine(positions, xo)
	xLine(positions, xo)
	hLine(positions, xo)
}

func hLine(positions []int, xo []string) {
	for i := range 15 {
		reserved := slices.Contains(positions, i)

		if reserved {
			fmt.Printf("%v", xo[rand.Intn(2)])
		}
		if i == 5 {
			fmt.Print("|")
		}

		if i == 10 {
			fmt.Print(" |")
		}

		if i == 5 || i == 10 && !reserved {
			// fmt.Print("|")
		} else if !reserved {
			fmt.Print(" ")
		}

		if i == 14 {
			fmt.Println("")
		}
	}
}

func xLine(positions []int, xo []string) {
	for i := range 15 {
		if i == 5 {
			fmt.Print("|")
		}

		if i == 10 {
			fmt.Print("|")
		}

		if i != 5 || i != 10 {
			fmt.Print("-")
		}

		if i == 14 {
			fmt.Println("")
		}
	}
}
