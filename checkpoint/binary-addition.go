package main

import "fmt"

func BinaryAddition(a int, b int) [32]int {
    sum := a + b
	i := 31
	final := [32]int{}
	for sum > 0 && i >= 0{
		index := sum % 2
		//final = append([]int{index},final...)
		final[i] = index 
		i--
		sum = sum / 2
	}
	return final
}
func main(){
	fmt.Println(BinaryAddition(1, 2176))
	fmt.Println(BinaryAddition(1, 2))
	fmt.Println(BinaryAddition(1, 3))
	fmt.Println(BinaryAddition(2, 1))
	fmt.Println(BinaryAddition(2, 2))
	fmt.Println(BinaryAddition(1, 16))
}