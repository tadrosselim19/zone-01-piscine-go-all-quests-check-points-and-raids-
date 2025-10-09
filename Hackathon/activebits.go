package main

import (
	"fmt"
)

func ActiveBits(n int) int {
	count := 0
	for n >= 1 {
		
		if n%2 == 1{
			count++
		}
		n =n/2
	}
	return count
}
func main() {
	fmt.Println(ActiveBits(7))
}