package main

import (
	"fmt"
)

func Compare(a, b string) int {
	arunes := []rune(a)
	brunes := []rune(b)

	minLen := min(len(arunes), len(brunes))

	for i := 0; i < minLen; i++ {
		if arunes[i] < brunes[i] {
			return -1
		} else if arunes[i] > brunes[i] {
			return 1
		}
	}

	// All compared runes are equal — now check lengths
	if len(arunes) < len(brunes) {
		return -1
	} else if len(arunes) > len(brunes) {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func aCompare(a, b string) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			if a[i] < b[i] {
				return -1
			}
			return 1
		}
	}
	if len(a) == len(b) {
		return 0
	}
	if len(a) < len(b) {
		return -1
	}
	return 1
}


func main() {
	fmt.Println(Compare("Salut!", "lut!"))   // -1
	fmt.Println(Compare("Hello!", "Hello!")) // 0
	fmt.Println(Compare("Ola!", "Ol"))       // 1
	fmt.Println(Compare("Hello!", "Hello"))
}
