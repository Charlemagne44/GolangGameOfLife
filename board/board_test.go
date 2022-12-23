package game

import (
	"testing"
)

func EquivalentSlice(initState, expectedState [][]int) bool {
	for i := 0; i < len(initState); i++ {
		for j := 0; j < len(initState[0]); j++ {
			if initState[i][j] != expectedState[i][j] {
				return false
			}
		}
	}
	return true
}

func TestNextStateDeadCell(t *testing.T) {
	// dead cell with no live neighbor
	initBoard := Board{
		Width:  3,
		Height: 3,
		Board: [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
	expectedBoard := Board{
		Width:  3,
		Height: 3,
		Board: [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
	initBoard.NextState()
	if !EquivalentSlice(initBoard.Board, expectedBoard.Board) {
		t.Fatal()
	}

	// one dead cell with 3 neighbors
	initBoard = Board{
		Width:  3,
		Height: 3,
		Board: [][]int{
			{1, 1, 1},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
	expectedBoard = Board{
		Width:  3,
		Height: 3,
		Board: [][]int{
			{0, 1, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
	}
	initBoard.NextState()
	if !EquivalentSlice(initBoard.Board, expectedBoard.Board) {
		t.Fatal()
	}
}

func TestNextStateAliveCell(t *testing.T) {
	// alive cell with 1 neighbor
	initBoard := Board{
		Width:  3,
		Height: 3,
		Board: [][]int{
			{1, 1, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
	expectedBoard := Board{
		Width:  3,
		Height: 3,
		Board: [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
	initBoard.NextState()
	if !EquivalentSlice(initBoard.Board, expectedBoard.Board) {
		t.Fatal()
	}
	// alive cell with 2 to 3 neighbors
	initBoard = Board{
		Width:  3,
		Height: 3,
		Board: [][]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}
	expectedBoard = Board{
		Width:  3,
		Height: 3,
		Board: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
	}
	initBoard.NextState()
	if !EquivalentSlice(initBoard.Board, expectedBoard.Board) {
		t.Fatal()
	}
	// alive cell with more than 3 neighbors
	initBoard = Board{
		Width:  3,
		Height: 3,
		Board: [][]int{
			{0, 1, 0},
			{1, 1, 1},
			{0, 1, 0},
		},
	}
	expectedBoard = Board{
		Width:  3,
		Height: 3,
		Board: [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
	}
	initBoard.NextState()
	if !EquivalentSlice(initBoard.Board, expectedBoard.Board) {
		t.Fatal()
	}

}
