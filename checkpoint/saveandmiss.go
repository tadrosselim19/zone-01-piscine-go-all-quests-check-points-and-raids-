package main

import (
	"fmt"
)
func SaveAndMiss(arg string, num int) string {
	if num <= 0{
		return arg
	}
	final := ""
	count := 0
	for i:= 0 ; i < len(arg) ; i++{
		if count != num{
			final += string(rune(arg[i]))
			count++
		}else{
			i+=num-1
			count = 0
		}
	}
	return final
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}