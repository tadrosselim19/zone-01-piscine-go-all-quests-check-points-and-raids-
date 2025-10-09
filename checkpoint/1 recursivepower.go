package main

import (
	"fmt"
)
func RecursivePower(nb int, power int) int {
	final := 1
	if power < 0{
		return 0
	}
	if power == 0 {
		return 1
	}
	if power >= 1{
		final = nb *RecursivePower( nb, power-1) 
	}
	return final
}
func main() {
	fmt.Println(RecursivePower(4, 3))
}