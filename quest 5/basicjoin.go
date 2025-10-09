package main

import (
	"fmt"
	"strings"
	
)

func BasicJoin(elems []string) string {
	new_line := ""
	for i := 0 ; i < len(elems) ; i++{
		new_line += elems[i]
	}
	return new_line
}
func aBasicJoin(elems []string) string {
	var builder strings.Builder
	for _, str := range elems {
		builder.WriteString(str)
	}
	return builder.String()
}

func Join(strs []string, sep string) string {
	new_line := ""
	for i := 0 ; i < len(strs) ; i++{
		new_line += strs[i]
		if i < len(strs)-1{
			new_line += sep
		}
	}
	return new_line
}
func aJoin(strs []string, sep string) string {
	return strings.Join(strs, sep)
}
func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}