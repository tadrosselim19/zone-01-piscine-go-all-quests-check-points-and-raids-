package main

import (
	"fmt"
)
func JumpOver(str string) string {
	srunes := []rune(str)
	word := ""
	if len(srunes) < 3{
		return "\n"
	}
	for i := 2 ; i < len(srunes) ; {
		word += string(srunes[i])
		i +=3
	}
	
	return word + "\n"
}
func main() {
	fmt.Print(JumpOver("1010101010"))
	fmt.Print(JumpOver(""))
	fmt.Print(JumpOver("t w e l v e"))
	fmt.Print(JumpOver("12"))
}