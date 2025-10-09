package main

import (
	"fmt"
)
func WeAreUnique(str1 string, str2 string) int {
	// special condtion mangment
	if str1 == "" && str2 == ""{
		return -1
	}else if len(str1) == 0 && len(str2) >0{
		return len(str2)
	}else if len(str2) == 0 && len(str1) >0{
		return len(str1)
	}
    // join the two strings and count the letter odd
	count := 0
	all_str := str1 + str2
	for i:= 0 ; i < len(all_str) ; i++{
		found :=false
		for j := 0; j < len(all_str); j++ {
			if i != j && all_str[i] == all_str[j] {
				found = true
			} 
		}
		if found ==false{
			count++
		}
	}
	
	return count
}
func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
	fmt.Println(WeAreUnique("oooo", "oo"))
	fmt.Println(WeAreUnique("ab", "aba"))
	fmt.Println(WeAreUnique("abc", "aba"))



}