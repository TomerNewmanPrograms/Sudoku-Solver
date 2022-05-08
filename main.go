package main

import (
	"fmt"
)

func main() {
	board := [9][9]int{{7, 8, 0, 4, 0, 0, 1, 2, 0},
		{6, 0, 0, 0, 7, 5, 0, 0, 9},
		{0, 0, 0, 6, 0, 1, 0, 7, 8},
		{0, 0, 7, 0, 4, 0, 2, 6, 0},
		{0, 0, 1, 0, 5, 0, 9, 3, 0},
		{9, 0, 4, 0, 6, 0, 0, 0, 5},
		{0, 7, 0, 3, 0, 0, 0, 1, 2},
		{1, 2, 0, 0, 0, 7, 4, 0, 0},
		{0, 4, 9, 2, 0, 6, 0, 0, 7},
	}
	printBoard(board)
	if solveSodoku(board, 0, 0) == true {
		printBoard(board)
	} else {
		fmt.Print("Can not be solved")
	}
}

func printBoard(board [9][9]int) {
	var i, j int
	fmt.Print("______________________\n")
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if j%3 == 0 {
				fmt.Print("|")
			}
			fmt.Print(" ", board[i][j])
		}
		if (i+1)%3 == 0 {
			fmt.Println()
			fmt.Print("______________________")
		}
		fmt.Println()

	}

}

func isValid(board [9][9]int, j, i int) bool {
	var boxX, boxY, row, col, index int
	boxX = i / 3
	boxY = j / 3
	for row = boxX; row < boxX+3; row++ {
		for col = boxY; col < boxY+3; col++ {
			if board[row][col] == board[i][j] && (i != row || j != col) {
				return false
			}
		}
	}
	for index = 0; i < 9; i++ {
		if (board[i][index] == board[i][j] && j != index) || (board[index][j] == board[i][j] && i != index) {
			return false
		}
	}
	return true
}
func solveSodoku(board [9][9]int, j, i int) bool {
	if i == 9 || j == 9 {
		return true
	}
	fmt.Println("SDa")
	var newJ, newI int
	Digit := 1
	if j == 8 {
		newI = i + 1
		newJ = 0
	} else {
		newI = i
		newJ = j + 1
	}
	for ; Digit <= 9; Digit++ {
		board[i][j] = Digit
		if isValid(board, j, i) == true && solveSodoku(board, newJ, newI) == true {
			return true
		}
	}
	board[i][j] = 0
	return false
}
