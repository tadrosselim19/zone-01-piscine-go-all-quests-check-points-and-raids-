package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	printStr("--insert")
	printStr("  -i")
	printStr("\t This flag inserts the string into the string passed as argument.")
	printStr("--order")
	printStr("  -o")
	printStr("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func bubbleSort(runes []rune) {
	n := len(runes)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	insert := ""
	order := false
	input := ""

	for _, arg := range args {
		switch {
		case len(arg) >= 9 && arg[:9] == "--insert=":
			insert = arg[9:]
		case len(arg) >= 3 && arg[:3] == "-i=":
			insert = arg[3:]
		case arg == "--order" || arg == "-o":
			order = true
		default:
			input = arg
		}
	}

	result := input + insert

	if order {
		runes := []rune(result)
		bubbleSort(runes)
		result = string(runes)
	}

	printStr(result)
}
