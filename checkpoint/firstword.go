package main

import (
	"os"

	"github.com/01-edu/z01"
)
func isWhitespace1(c byte) bool {
    return c == ' ' || c == '\t'
}
func main() {
	if len(os.Args) != 2 {
		return
	}
	sentence := os.Args[1]
	if sentence == "" {
		return
	}
	// having the starting point
	stared_index := 0
	for i := 0 ; i < len(sentence) ; i++ {
		if isWhitespace1(sentence[i]) == false{
			break
		}
		if isWhitespace1(sentence[i]){
			stared_index++
		}
	}
	// starting print frist word by its index
	for i := stared_index ; i < len(sentence) ; i++{
		if sentence[i] == ' ' || sentence[i] == '\t'{
			break
		}else{
			z01.PrintRune(rune(sentence[i]))
		}
	}
	z01.PrintRune('\n')
}