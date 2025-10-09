package main

import (
	"fmt"
)
func tIsSorted(f func(a, b int) int, a []int) bool {
	result := false
	for i := 0 ; i+1 < len(a) ; i++{
		if f(a[i] ,a[i+1]) == 1 {
			result = false
			break
		}
		result = true
		
	}

	return result
}
func f(a,b int)int{
	if a > b {
		return 1
	}else if a == b{
		return 0
	}
	return -1
}
func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	ascending := true
	descending := true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			ascending = false
		}
		if f(a[i], a[i+1]) < 0 {
			descending = false
		}
	}

	return ascending || descending
}
func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{5, 4, 3}
	a4 := []int{1, 1, 1}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(IsSorted(f, a3))
	fmt.Println(IsSorted(f, a4))
}