package main

import "fmt"

func birthday(s []int32, d int32, m int32) int32 {
	var count int32 = 0
	length := 0
	sum := 0

	for i := 0; i < len(s); i++ {
		length = 0
		sum = 0
		for length != int(m) && i+length <len(s){
			sum += int(s[i+length])
			length++
		}
		
		if sum == int(d) {
			count++
		}
		

	}
	return count
}
func main() {
	fmt.Println(birthday([]int32{1,1,1,1,1,1},2,3))
	fmt.Println(birthday([]int32{1,2,1,3,2,1},3,2))
	fmt.Println(birthday([]int32{4},4,1))



}