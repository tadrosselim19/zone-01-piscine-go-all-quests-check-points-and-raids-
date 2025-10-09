package main

import (
	"fmt"
)
func aIndex(s string, toFind string) int {
	srunes := []rune(s)
	toFindrunes := []rune(toFind)
		for j := 0 ; j < len(srunes) ; j++{
			for i:= 0 ; i < len(toFind) ; i++{
			   if toFindrunes[0] == srunes[j]{
				if toFindrunes[i] == srunes[j]{
				   return j
			   }   
			}
			}

		}
	return -1
// Index("fq9alpM}]'s.}", "") == -1 instead of 0
// exit status 1
}

func Index(s string, toFind string) int {
	srunes := []rune(s)
	toFindRunes := []rune(toFind)

	for i := 0; i <= len(srunes)-len(toFindRunes); i++ {
		match := true
		for j := 0; j < len(toFindRunes); j++ {
			if srunes[i+j] != toFindRunes[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}


func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
	fmt.Println(Index("Salut!", "Salut"))
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}


}