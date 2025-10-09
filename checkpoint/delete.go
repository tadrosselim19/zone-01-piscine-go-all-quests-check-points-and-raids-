package main

import "fmt"

func Delete(ints []int, position int) []int {
	final := []int{}
	for i := 0 ; i < len(ints) ; i++{
		if i+1 != position{
			final = append(final, ints[i])
		}
	}
	return final
}

func main() {
	fmt.Println(Delete([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(Delete([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(Delete([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(Delete([]int{1, 2, 3, 4, 5}, -1))

}