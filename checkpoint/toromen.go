package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	input := os.Args[1]
	n := 0

	// Convert input string to integer manually
	for i := 0; i < len(input); i++ {
		ch := input[i]
		if ch < '0' || ch > '9' {
			fmt.Println("ERROR: cannot convert to roman digit")
			return
		}
		n = n*10 + int(ch-'0')
	}

	if n < 1 || n > 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	fmt.Println(convertToRoman(n))
}

func convertToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			result += symbols[i]
		}
	}
	return result
}