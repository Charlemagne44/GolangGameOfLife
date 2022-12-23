package game

import (
	"bufio"
	"os"
)

// create board from give filename
func LoadFromFile(filename string) Board {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lines := getStringSlice(file)
	width := len(lines[0])
	height := len(lines)
	return createBoardFromSlice(width, height, lines)
}

// create board object from string slice
func createBoardFromSlice(width, height int, lines []string) Board {
	// create board object
	board := Board{
		Height: height,
		Width:  width,
	}
	// give board space in memory
	board.Board = make([][]int, height)
	for i := 0; i < height; i++ {
		board.Board[i] = make([]int, width)
	}
	// assign values from string array
	for i, line := range lines {
		for j, char := range line {
			board.Board[i][j] = int(char - '0')
		}
	}
	return board
}

// create slice of strings from input file
func getStringSlice(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
