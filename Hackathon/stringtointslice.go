package main

import (
	"fmt"
	"github.com/01-edu/z01"
)
func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}
	strrunes := []rune(str)
	result := []int{}

	for _, i := range strrunes {
		result = append(result, int(i))
	}

	return result
}

func main() {
	fmt.Println(StringToIntSlice("A quick brown fox jumps over the lazy dog"))
	fmt.Println(StringToIntSlice("Converted this string into an int"))
	fmt.Println(StringToIntSlice("hello THERE"))
}