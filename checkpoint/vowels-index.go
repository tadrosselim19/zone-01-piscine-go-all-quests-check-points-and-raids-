package main

import "fmt"
func VowelsIndex(str string) []int {
	index := []int{}
	vowels := []rune{'a','e','i','o','u','A','E','I','O','U',}
	r_str := []rune(str)
	for i := 0 ; i < len(r_str) ; i++{
		for j := 0 ; j < len(vowels) ; j++{
		   if r_str[i] == vowels[j]{
			index = append(index, i)
			break
		   }
		}
	}
	return index
}
func main() {
	res := []string{"student", "hello Iyan","bcdfgh", "wOrld", "","AAEO$o;jw"}
	for _, i := range res {
		fmt.Println(VowelsIndex(i))
	}
}