package game

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
func (board *Board) InitDeadBoard() {
	// initialize  the board
	board.Board = make([][]int, board.Height)
	for i := 0; i < board.Height; i++ {
		board.Board[i] = make([]int, board.Width)
	}
}

// print the board in matrix format
func (board *Board) PrintBoard() {
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
func (board *Board) RandomizeBoard(likelyAlive float64) {
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

// this functionc calculates the next board state
// in the game of life. 1 alive neighbor to alive -> dead.
// 2-3 -> alive, > 3 -> dead, dead cell w/ 3 alive
// neighbors -> alive
func (board *Board) NextState() {
	for y, row := range board.Board {
		for x, col := range row {
			alive_count := 0
			if y == 0 && x == 0 {
				// top left corner
				alive_count = board.Board[y+1][x] + // s
					board.Board[y][x+1] + // e
					board.Board[y+1][x+1] // se
			} else if y == 0 && x == board.Width-1 {
				// top right corner
				alive_count = board.Board[y+1][x] + // s
					board.Board[y][x-1] + // w
					board.Board[y+1][x-1] // sw
			} else if y == board.Height-1 && x == 0 {
				// bottom left corner
				alive_count = board.Board[y-1][x] + // n
					board.Board[y][x+1] + // e
					board.Board[y-1][x+1] // ne
			} else if y == board.Height-1 && x == board.Height-1 {
				// bottom right corner
				alive_count = board.Board[y-1][x] + // n
					board.Board[y][x-1] + // w
					board.Board[y-1][x-1] // nw
			} else if y == 0 {
				// top edge
				alive_count = board.Board[y][x-1] + // w
					board.Board[y+1][x-1] + // sw
					board.Board[y+1][x] + // s
					board.Board[y+1][x+1] + // se
					board.Board[y][x+1] // e
			} else if y == board.Height-1 {
				// bottom edge
				alive_count = board.Board[y][x-1] + // w
					board.Board[y-1][x-1] + // nw
					board.Board[y-1][x] + // n
					board.Board[y-1][x+1] + // ne
					board.Board[y][x+1] // e
			} else if x == 0 {
				// left edge
				alive_count = board.Board[y-1][x] + // n
					board.Board[y-1][x+1] + // ne
					board.Board[y][x+1] + // e
					board.Board[y+1][x+1] + // se
					board.Board[y+1][x] // s
			} else if x == board.Width-1 {
				// right edge
				alive_count = board.Board[y-1][x] + // n
					board.Board[y-1][x-1] + // nw
					board.Board[y][x-1] + // w
					board.Board[y+1][x-1] + // sw
					board.Board[y+1][x] // s
			} else {
				// center
				alive_count = board.Board[y-1][x] + // n
					board.Board[y-1][x-1] + // nw
					board.Board[y][x-1] + // w
					board.Board[y+1][x-1] + // sw
					board.Board[y+1][x] + // s
					board.Board[y+1][x+1] + // se
					board.Board[y][x+1] + // e
					board.Board[y-1][x+1] // ne
			}

			// reassign values based on neighor counts
			if col == 1 {
				if alive_count <= 1 {
					board.Board[y][x] = 0
				} else if alive_count > 3 {
					board.Board[y][x] = 0
				}
			} else {
				if alive_count == 3 {
					board.Board[y][x] = 1
				}
			}
		}
	}
}
