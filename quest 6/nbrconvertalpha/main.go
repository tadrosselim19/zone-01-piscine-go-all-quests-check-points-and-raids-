package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	// Check for --upper flag
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	printed := false // track if any character was printed

	for _, arg := range args {
		n := 0
		valid := true

		for _, r := range arg {
			if r < '0' || r > '9' {
				valid = false
				break
			}
			n = n*10 + int(r-'0')
		}

		if !valid || n < 1 || n > 26 {
			z01.PrintRune(' ')
			printed = true
			continue
		}

		var letter rune
		if upper {
			letter = rune('A' + n - 1)
		} else {
			letter = rune('a' + n - 1)
		}
		z01.PrintRune(letter)
		printed = true
	}

	if printed {
		z01.PrintRune('\n')
	}
}
