package main

import (
	"fmt"
	
)
func Join(strs []string, sep string) string {
	result := ""
	for i := 0 ; i < len(strs) ; i++{
		result = result + strs[i]
		if i != len(strs)-1{
			result += sep
		}
	}
	return result
}
func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}