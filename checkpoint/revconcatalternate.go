package main

import (
	"fmt"
)
func RevConcatAlternate(slice1,slice2 []int) []int {
	sl1 := len(slice1)
	sl2 := len(slice2)
	largest := 0	
	final := make([]int,len(slice1)+ len(slice2))
	for sl1 != sl2{
	   if sl1 > sl2{
		largest = sl1 - sl2
		for i:=0 ; sl1 > sl2 ; i++{
			final[i] = slice1[sl1-1]
			sl1--
		}
       }
	   if sl1 < sl2{
		largest = sl2 - sl1
		for i:=0 ; sl2 >sl1; i++{
			final[i] = slice2[sl2-1]
			sl2--
		}
	   }
    }
	for i:= largest; i < len(final);{
		if sl1 > 0 && sl2 >0 {
			final[i] = slice1[sl1-1]
			sl1--
			i++
			final[i] = slice2[sl2-1]
			sl2--
			i++
		}else{
			break
		}
	}
	return final


}
func RevConcatAlternate1(slice1, slice2 []int) []int {
    res := []int{}
    i, j := len(slice1)-1, len(slice2)-1

    if len(slice1) > len(slice2) {
        for ; i >= len(slice2); i-- {
            res = append(res, slice1[i])
        }
    } else if len(slice2) > len(slice1) {
        for ; j >= len(slice1); j-- {
            res = append(res, slice2[j])
        }
    }

    for i >= 0 && j >= 0 {
        res = append(res, slice1[i])
        i--
        res = append(res, slice2[j])
        j--
    }

    return res
}


func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}