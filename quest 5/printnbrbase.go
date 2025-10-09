package main

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := len(base)
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	num := int64(nbr)
	if num < 0 {
		z01.PrintRune('-')
		num = -num
	}

	printInBase(uint64(num), base, baseLen)
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, ch := range base {
		if ch == '+' || ch == '-' || seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

func printInBase(n uint64, base string, baseLen int) {
	if n >= uint64(baseLen) {
		printInBase(n/uint64(baseLen), base, baseLen)
	}
	z01.PrintRune(rune(base[n%uint64(baseLen)]))
}
func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}