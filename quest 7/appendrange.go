package main

import (
	"fmt"
)

func main() {
	// fmt.Println(AppendRange(5, 10))
	// fmt.Println(AppendRange(10, 5))
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	s := []int{}
	for i := min ; i < max ; i++{
		s = append(s, i)
	}
	return s

}
func aMakeRange(min, max int) []int {
	if min >= max {
		return nil
	} 
	s := make([]int, max-min)
	for j := 0 ; j < len(s) ; j++{
	    for i := min ; i < max ; i++{
			s[j]=i
			min = min +1
			break
		}
		
	}
	return s
}
func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	} 
	s := make([]int, max-min)
	for i := range(s){
		s[i]=min
		min += 1
	}
	return s
}
