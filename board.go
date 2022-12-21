package main

import (
	"fmt"
	"math/rand"
)

type Board struct {
	Width  int
	Height int
	Board  [][]int
}

// create a board of a certain width and height where
// all values are left as 0 (dead)
func (board *Board) initDeadBoard() {
	// initialize  the board
	board.Board = make([][]int, board.Height)
	for i := 0; i < board.Height; i++ {
		board.Board[i] = make([]int, board.Width)
	}
}

// print the board in matrix format
func (board *Board) printBoard() {
	for _, row := range board.Board {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Printf("\n")
	}
}

// randomize the values of the board between
// 0 and 1 (dead and alive), likely alive is
// the factor used in RNG, think of it as the
// % chance we set a particular cell to 1(alive)
func (board *Board) randomizeBoard(likelyAlive float64) {
	for y := 0; y < len(board.Board); y++ {
		for x := 0; x < len(board.Board[0]); x++ {
			if rand.Float64() < likelyAlive {
				board.Board[x][y] = 1
			} else {
				board.Board[x][y] = 0
			}
		}
	}
}
