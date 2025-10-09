package main

import (
	"fmt"
	"os"
)

// Validate input
func isValidInput(args []string) bool {
	if len(args) != 9 {
		return false
	}
	for _, line := range args {
		if len(line) != 9 {
			return false
		}
		for _, ch := range line {
			if (ch < '1' || ch > '9') && ch != '.' {
				return false
			}
		}
	}
	return true
}

// Convert input to board
func makeBoard(args []string) [][]byte {
	board := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		board[i] = make([]byte, 9)
		for j := 0; j < 9; j++ {
			board[i][j] = args[i][j]
		}
	}
	return board
}

// Check if placing digit is valid
func isValid(board [][]byte, row, col int, digit byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == digit || board[i][col] == digit {
			return false
		}
	}
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == digit {
				return false
			}
		}
	}
	return true
}

// Backtracking solver
func solve(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for d := byte('1'); d <= '9'; d++ {
					if isValid(board, row, col, d) {
						board[row][col] = d
						if solve(board) {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

// Print the solved board
func printBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", board[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Main
func main() {
	args := os.Args[1:]

	if !isValidInput(args) {
		fmt.Println("Error")
		return
	}

	board := makeBoard(args)

	if !solve(board) {
		fmt.Println("Error")
		return
	}

	printBoard(board)
}
