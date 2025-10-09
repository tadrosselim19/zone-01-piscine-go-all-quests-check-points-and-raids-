package main

import (
	"fmt"
)
func FifthAndSkip(str string) string {
	final := ""
	if str == ""{
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	count := 0
	for i:= 0 ; i < len(str) ; i++{
		if str[i] == ' '{
			continue
		}
		if count == 5 {
			final += " "
			count = 0
		}else {
			final += string(rune(str[i]))
			count++
		}
	}
	return final + "\n"
}
func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}