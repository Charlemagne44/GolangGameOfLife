package main

import "fmt"

func (board *Board) render() {
	// top edge
	for i := 0; i < (board.Width*3)+2; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
	// board contents + side edges
	for _, row := range board.Board {
		fmt.Print("|")
		for _, col := range row {
			if col == 0 {
				fmt.Print("   ")
			} else {
				fmt.Print(" # ")
			}
		}
		fmt.Print("|\n")
	}

	// bottom edge
	for i := 0; i < (board.Width*3)+2; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}
