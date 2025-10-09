package main

import (
	"fmt"
)
func ConcatAlternate(slice1,slice2 []int) []int {
	slice1_len := len(slice1) 
	slice2_len := len(slice2)
	slice3_len := len(slice1) 
	slice4_len := len(slice2)
	sum := slice1_len + slice2_len
// detection if any of sclices are empty
	if slice1_len == 0 && slice2_len == 0 {
		return nil
	}else if slice1_len == 0{
		return slice2
	}else if slice2_len == 0{
		return slice1
	}
// detection the largest sclice 
	subtraction,flag := 0 , 0
	if slice2_len >= slice1_len{
		subtraction += slice2_len - slice1_len
		flag += 2 
	}else {
		 subtraction += slice1_len - slice2_len
		 flag ++
	}
// sclice inationation according to given sclices 
	final := make([]int,sum)

// printing the larest 
	if subtraction != 0 && flag == 1{
		for i:= sum-1 ; i >= subtraction && subtraction !=0  ;i--{
			final[i] = slice1[slice1_len-1]
			slice1_len--
			sum--
			subtraction--
		}
	}else if subtraction != 0 && flag == 2{
		for i:= sum-1 ; i >= subtraction && subtraction !=0  ;i--{
			final[i] = slice2[slice2_len-1]
			slice2_len--
			sum--
			subtraction--
		}
	}	
// alternative the rest
	for i:= sum-1 ; i >= 0  ;i-=2{
		if slice4_len > slice3_len{
			final[i-1] = slice2[slice2_len-1]
			final[i] =slice1[slice1_len-1]
			slice1_len--
			slice2_len--
		}else{
			final[i-1] =slice1[slice1_len-1]
			final[i] = slice2[slice2_len-1]
			slice1_len--
			slice2_len--
		}
	}
	


	return final
}
func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}


package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	// If one slice is empty, return the other
	if len(slice1) == 0 {
		return slice2
	}
	if len(slice2) == 0 {
		return slice1
	}

	var first, second []int
	if len(slice1) >= len(slice2) {
		first, second = slice1, slice2
	} else {
		first, second = slice2, slice1
	}

	result := make([]int, 0, len(slice1)+len(slice2))
	i := 0
	// Alternate elements
	for i < len(second) {
		result = append(result, first[i], second[i])
		i++
	}
	// Append remaining elements from the longer slice
	result = append(result, first[i:]...)

	return result
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))        // [1 4 2 5 3 6]
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11})) // [1 2 3 4 5 6 7 8 9 10 11]
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))       // [4 1 5 2 6 3 7 8 9]
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))                         // [1 2 3]
}
