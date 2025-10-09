package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1] + os.Args[2]
	seen := make(map[rune]bool)
	result := []rune{}

	for _, ch := range input {
		if !seen[ch] {
			seen[ch] = true
			result = append(result, ch)
		}
	}

	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
