package main

import "fmt"
func TwoSum(nums []int, target int) []int {
	for i := 0 ; i < len(nums) ;i++{
		for j := 0 ; j < len(nums) ;j++{
			if i != j && nums[i] + nums[j]  == target{
				return []int{i,j}
			}
		}
		 
	}
	return nil
}
func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}