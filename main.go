package main

import "fmt"

func main() {
	var board [9][9]int
	if solveSodoku(board) == true {
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
func solveSodoku(board [9][9]int) bool {

}
