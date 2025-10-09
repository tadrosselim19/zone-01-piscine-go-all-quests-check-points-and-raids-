package main

import (
	"fmt"

)

func Chunk(slice []int, size int) {
	// validation
	if size <= 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		fmt.Println("[]")
		return
	}
	//variables initation
	large_array := [][]int{}
	frist_array := []int{}
	modified_size := len(slice)
	new_size := size
    stated := 0
for modified_size > 0{
	for i:=stated ; i<len(slice) ;i++{
		if i < new_size {
			frist_array = append(frist_array,slice[i])
		}	
	}
	large_array = append(large_array, frist_array)
	modified_size -= size
	stated += size
	new_size+= size
	frist_array = []int{}
}
	fmt.Println(large_array)
}
func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}