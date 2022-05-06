package main

import "fmt"

func main() {
	var board [9][9]int
	if solveSodoku(board, 0, 0) == true {
		printBoard(board)
	}
}

func printBoard(board [9][9]int) {
	var i, j int
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			fmt.Print(" %d ", board[i][j])
		}
		fmt.Println()

	}

}
func isValid(board [9][9]int, j, i int) bool {
	// to do - return if the new assignment is valid
}
func solveSodoku(board [9][9]int, j, i int) bool {
	// to do - recursion
}
