package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Join args with spaces
	var runes []rune
	for i, arg := range args {
		for _, r := range arg {
			runes = append(runes, r)
		}
		if i != len(args)-1 {
			runes = append(runes, ' ')
		}
	}

	// Collect vowels positions and runes
	var vowelsPos []int
	var vowelsRunes []rune
	for i, r := range runes {
		if isVowel(r) {
			vowelsPos = append(vowelsPos, i)
			vowelsRunes = append(vowelsRunes, r)
		}
	}

	if len(vowelsPos) == 0 {
		// No vowels, print as is
		for _, r := range runes {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
		return
	}

	// Reverse vowels runes
	for i, j := 0, len(vowelsRunes)-1; i < j; i, j = i+1, j-1 {
		vowelsRunes[i], vowelsRunes[j] = vowelsRunes[j], vowelsRunes[i]
	}

	// Replace vowels in original string with reversed vowels
	for i, pos := range vowelsPos {
		runes[pos] = vowelsRunes[i]
	}

	// Print final string
	for _, r := range runes {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
