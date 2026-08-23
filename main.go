package main

import (
	"fmt"
	"strings"
)

type Cell byte

const (
	Empty Cell = iota
	X
	O
)

func (c Cell) Value() byte {
	switch c {
	case X:
		return 'X'
	case O:
		return 'O'
	default:
		return ' '
	}
}

type board struct {
	cells   [9]Cell
	current Cell
}

func NewBoard() *board {
	return &board{current: X}
}

func (b *board) getValue(pos int) Cell {
	return b.cells[pos]
}

func (b *board) incrementCurrent() {
	if b.current == X {
		b.current = O
	} else {
		b.current = X
	}
}

var winLines = [8][3]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

var rows = [3][3]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
}

func (b *board) draw() string {
	var sb strings.Builder
	pad := [2]byte{' ', ' '}

	sb.WriteByte('\n')
	for rowIdx, row := range rows {
		for colIdx, pos := range row {
			for _, p := range pad {
				sb.WriteByte(p)
			}
			sb.WriteByte(b.cells[pos].Value())
			for _, p := range pad {
				sb.WriteByte(p)
			}
			if colIdx < len(row)-1 {
				sb.WriteByte('|')
			}
		}
		sb.WriteByte('\n')

		if rowIdx < len(rows)-1 {
			for i := range 3 {
				for range 5 {
					sb.WriteByte('-')
				}
				if i < 2 {
					sb.WriteByte('|')
				}
			}
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}

func main() {
	b := NewBoard()
	fmt.Print(b.draw())
}
