package main

import (
	"fmt"
)

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for i := 0 ; i < len(tab) ; i++{
		if f(tab[i]) == true{
			count ++
		}
	}
	return count 

}
func IsNumeric(s string)bool{
	for _, i := range s{
		if i >= '0' && i <='9'{
			return true
		}
	}
	return false
}

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This","1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}