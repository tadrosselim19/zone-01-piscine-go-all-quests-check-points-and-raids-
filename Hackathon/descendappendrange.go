package main

import (
	"fmt"
	
)
func DescendAppendRange(max, min int) []int {
	result := []int{}
	if max < min{
		return []int{}
	}
	for i := max ; i > min ; i--{
		result= append(result, i)
	}
	return result
}
func main() {
	fmt.Println(DescendAppendRange(10, 5))
	fmt.Println(DescendAppendRange(5, 10))
}