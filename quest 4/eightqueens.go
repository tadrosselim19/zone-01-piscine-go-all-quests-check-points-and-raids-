package main

import (
	"fmt"
)

const size = 8

func main() {
	solve([]int{})
}

// solve attempts to place queens on the board using backtracking.
func solve(queens []int) {
	if len(queens) == size {
		printSolution(queens)
		return
	}

	for row := 1; row <= size; row++ {
		if isSafe(queens, row) {
			solve(append(queens, row))
		}
	}
}

// isSafe checks if placing a queen in the next column at 'row' is safe.
func isSafe(queens []int, row int) bool {
	col := len(queens)
	for i, qRow := range queens {
		if qRow == row ||                // same row
			qRow-(col-i) == row ||       // same \ diagonal
			qRow+(col-i) == row {        // same / diagonal
			return false
		}
	}
	return true
}

// printSolution prints the current arrangement as a single line string.
func printSolution(queens []int) {
	for _, row := range queens {
		fmt.Print(row)
	}
	fmt.Println()
}
