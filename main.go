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
	printBoard(&board)
	if solveSodoku(&board, 0, 0) == true {
		printBoard(&board)
	} else {
		fmt.Print("Can not be solved")
	}
}

func printBoard(board *[9][9]int) {
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

func isValid(board *[9][9]int, j, i int) bool {
	var boxX, boxY, row, col, index int
	if j == 2 && i == 0 && board[i][j] == 3 {
		index = 0
	}
	boxX = i / 3
	boxY = j / 3
	for row = boxX * 3; row < 3*boxX+3; row++ {
		for col = boxY * 3; col < 3*boxY+3; col++ {
			if board[row][col] == board[i][j] && (i != row || j != col) {
				return false
			}
		}
	}
	for index = 0; index < 9; index++ {
		if (board[i][index] == board[i][j] && j != index) || (board[index][j] == board[i][j] && i != index) {
			return false
		}
	}
	return true
}
func solveSodoku(board *[9][9]int, j, i int) bool {
	if i == 9 || j == 9 {
		return true
	}
	var newJ, newI int
	Digit := 1
	currJ := j
	currI := i
	for currI < 9 {
		if board[currI][currJ] == 0 {
			break
		}
		if currJ == 8 {
			currI = currI + 1
			currJ = 0
		} else {
			currJ = currJ + 1
		}
	}
	if currJ == 8 {
		newI = currI + 1
		newJ = 0
	} else {
		newI = currI
		newJ = currJ + 1
	}
	if currI == 9 {
		return true
	}
	if board[0][2] == 5 && board[0][4] == 3 && currJ == 8 {
		board[0][2] = 5
	}
	for ; Digit <= 9; Digit++ {
		board[currI][currJ] = Digit
		if isValid(board, currJ, currI) == true && solveSodoku(board, newJ, newI) == true {
			return true
		}
	}
	board[currI][currJ] = 0
	return false
}
