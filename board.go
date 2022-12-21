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
func initDeadBoard(width, height int) [][]int {
	// initialize  the board
	deadBoard := make([][]int, height)
	for i := 0; i < height; i++ {
		deadBoard[i] = make([]int, width)
	}
	return deadBoard
}

// print the board in matrix format
func printBoard(board [][]int) {
	for _, row := range board {
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
func randomizeBoard(board [][]int, likelyAlive float64) {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if rand.Float64() < likelyAlive {
				board[x][y] = 1
			} else {
				board[x][y] = 0
			}
		}
	}
}
