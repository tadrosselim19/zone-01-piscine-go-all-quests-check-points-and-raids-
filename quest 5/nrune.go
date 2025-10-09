package main

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	srunes := []rune(s)
	if n < 1 || n > len(srunes) {
		return 0
	}
	return srunes[n-1]
}
func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}