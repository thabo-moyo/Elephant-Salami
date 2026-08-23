package main

import "testing"

func TestNewBoardDefaults(t *testing.T) {
	b := NewBoard()
	if b.current != X {
		t.Errorf("expected current player X, got %v", b.current)
	}
	for i, c := range b.cells {
		if c != Empty {
			t.Errorf("expected cell %d to be Empty, got %v", i, c)
		}
	}
}

func TestGetValue(t *testing.T) {
	b := NewBoard()
	b.cells[4] = O
	if got := b.getValue(4); got != O {
		t.Errorf("expected O at position 4, got %v", got)
	}
}

func TestIncrementCurrent(t *testing.T) {
	b := NewBoard()
	if b.current != X {
		t.Fatalf("expected initial player X, got %v", b.current)
	}
	b.incrementCurrent()
	if b.current != O {
		t.Errorf("expected current player O after increment, got %v", b.current)
	}
	b.incrementCurrent()
	if b.current != X {
		t.Errorf("expected current player X after second increment, got %v", b.current)
	}
}

func TestCellValue(t *testing.T) {
	cases := []struct {
		cell Cell
		want byte
	}{
		{X, 'X'},
		{O, 'O'},
		{Empty, ' '},
	}
	for _, tc := range cases {
		if got := tc.cell.Value(); got != tc.want {
			t.Errorf("Cell(%v).Value() = %q, want %q", tc.cell, got, tc.want)
		}
	}
}

func TestDrawContainsPlacedValues(t *testing.T) {
	b := NewBoard()
	b.cells[0] = X
	b.cells[4] = O
	out := b.draw()

	if !containsByte(out, 'X') {
		t.Error("expected drawn board to contain 'X'")
	}
	if !containsByte(out, 'O') {
		t.Error("expected drawn board to contain 'O'")
	}
}

func containsByte(s string, b byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return true
		}
	}
	return false
}
