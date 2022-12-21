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
// 0 and 1 (dead and alive)
func randomizeBoard(board [][]int) {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			board[y][x] = rand.Intn(2)
		}
	}
}
