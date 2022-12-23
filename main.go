package main

import (
	"flag"
	game "gameOfLife/board"
	"time"
)

/*
1.Build a data structure to store the board state
2.“Pretty-print” the board to the terminal
3.Given a starting board state, calculate the next one
4.Run the game forever
5. Save interesting starting positions to files and add the ability to reload them into your Life
6. Improve the User Interface
7. Change the rules of Life
*/

func main() {
	height := flag.Int("h", 10, "height of the board")
	width := flag.Int("w", 10, "width of the board")
	turns := flag.Int("t", 10, "how many turns of the game to play")
	sleep := flag.Int("s", 3, "seconds of sleep between each turn")
	probability := flag.Float64("p", 0.50, "probability starting cell is alive")
	file := flag.String("f", "", "load a premade board from a file")
	flag.Parse()

	var board game.Board
	if *file != "" {
		board = game.LoadFromFile(*file)
	} else {
		board = game.Board{
			Width:  *width,
			Height: *height,
		}
		board.InitDeadBoard()
		board.RandomizeBoard(*probability)
	}
	for i := 0; i < *turns; i++ {
		board.Render()
		board.NextState()
		time.Sleep(time.Second * time.Duration(*sleep))
	}
}
