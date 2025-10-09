package main

import (
	"fmt"
)
func is_repeted_in_in_sclice(sl []int ,num int )bool{
	for i := 0 ; i < len(sl) ; i++{
		if  sl[i] == num{
			return true
		} 
	}
	return false
}
func RemoveDuplicate(slice []int) []int {
	final := []int{}
	for i := 0 ; i < len(slice) ; i++{
		if is_repeted_in_in_sclice(final , slice[i])==false{
			final = append(final, slice[i])
		}
	}
	return final
}
func RemoveDuplicate1(slice []int) []int {
    seen := make(map[int]bool)
    final := []int{}
    for _, num := range slice {
        if !seen[num] {
            seen[num] = true
            final = append(final, num)
        }
    }
    return final
}

func main() {
	fmt.Println(RemoveDuplicate([]int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(RemoveDuplicate([]int{1, 1, 2, 2, 3}))
	fmt.Println(RemoveDuplicate([]int{}))
}
