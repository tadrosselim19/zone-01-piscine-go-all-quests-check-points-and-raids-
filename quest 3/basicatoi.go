package main

import (
	"fmt"
	
)
func BasicAtoi2(s string) int {
	result := 0
	for i := 0 ; i <len(s) ; i++{
		if s[i] > '9' || s[i] < '0'{
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}
	return result
}
func BasicAtoi(s string) int {
	result := 0
	for i := 0 ; i <len(s) ; i++{
		result = result*10 + int(s[i]-'0')
	}
	return result
}
func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if table[i] < table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
func aortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
	s := []int{5,4,3,2,1,0}
	SortIntegerTable(s)
	fmt.Println(s)
	
}