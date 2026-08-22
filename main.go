package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

var (
	vSpace   = byte(' ')
	vX       = byte('X')
	vO       = byte('O')
	vHLine   = byte('-')
	vVLine   = byte('|')
	vNewLine = byte('\n')
)

type (
	board struct {
		cols  [5]int
		rows  [15]int
		state boardState
	}

	boardState struct {
		current int
	}
)

func main() {
	// basic 3x3 board
	for range 2 {
		fmt.Println(" ")
	}
	b := &board{}
	fmt.Printf("%v \n", b.draw())

	/**
	  X  |  X  |  X
	-----|-----|-----
	  O  |  X  |  O
	-----|-----|-----
	  O  |  X  |  O
	*/
	// positions := []int{2, 8, 13}

	// xo := []string{"X", "O"}
	// b := board{}
	// b.hLine(positions, xo)
	// xLine(positions, xo)
	// b.hLine(positions, xo)
	// xLine(positions, xo)
	// b.hLine(positions, xo)
}

func (b *board) draw() string {
	var data strings.Builder

	for range b.cols {
		for range b.rows {
			data.WriteByte(vO)
		}
		data.WriteByte(vNewLine)
	}

	return data.String()
}

func (b *board) current() byte {
	return vO
}

func (s *board) hLine(positions []int, xo []string) {
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
