package main

import (
	"fmt"
	
)
func Abort(a, b, c, d, e int) int {
	scliceofargument := []int{a,b,c,d,e}
	for i:= 0 ; i < len(scliceofargument) ; i++{
		for j:= 0 ; j < len(scliceofargument) ; j++{
			if scliceofargument[i] > scliceofargument[j]{
				scliceofargument[i] , scliceofargument[j] = scliceofargument[j] , scliceofargument[i]
			}
		}
	}
	return scliceofargument[len(scliceofargument)/2]
}
func bAbort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}
	for i := 0; i < len(nums); i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums[2]
}

func main() {
	middle := Abort(100, 10, 50, 20, 30)
	fmt.Println(middle)
}