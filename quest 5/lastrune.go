package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	runes := []rune(s)
	return runes[0]
}
func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}