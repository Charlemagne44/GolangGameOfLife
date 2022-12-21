package main

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
	board := &Board{
		Width:  10,
		Height: 10,
	}
	board.initDeadBoard()
	board.randomizeBoard(.20)
	board.printBoard()
	board.render()
}
