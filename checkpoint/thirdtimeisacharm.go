package main

import (
	"fmt"
)
func ThirdTimeIsACharm(str string) string {
	str_rune := []rune(str)
	final := ""
	for i:= 2 ; i < len(str_rune) ; i+=3{
		final+=string(str_rune[i])
	}
	return final + "\n"
}
func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
	fmt.Println(78%5)
}