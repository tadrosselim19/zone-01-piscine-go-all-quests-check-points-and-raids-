package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	board := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		line := os.Args[i+1]
		if len(line) != 9 {
			fmt.Println("Error")
			return
		}
		board[i] = []byte(line)
	}

	if !isValidInitialBoard(board) {
		fmt.Println("Error")
		return
	}

	if !solveSudoku(board) {
		fmt.Println("Error")
		return
	}

	printBoard(board)
}

func printBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", board[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println("$")
	}
}

func solveSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for d := byte('1'); d <= '9'; d++ {
					if IsNumberPositionValid(board, i, j, d) {
						board[i][j] = d
						if solveSudoku(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func IsNumberPositionValid(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
		if board[i][col] == num {
			return false
		}
	}
	startRow := row - row%3
	startCol := col - col%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}
	return true
}

func isValidInitialBoard(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val != '.' {
				board[i][j] = '.'
				if !IsNumberPositionValid(board, i, j, val) {
					return false
				}
				board[i][j] = val
			}
		}
	}
	return true
}
