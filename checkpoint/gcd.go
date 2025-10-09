package main

import (
	"fmt"
	"os"
)
func convert3(num string)int{
	rune_num := []rune(num)
	final := 0	
	for i := 0 ;i < len(rune_num); i++{
		if rune_num[i] <= '9' && rune_num[i] >= '0'{
			final = final *10 + int(rune_num[i])-'0'
		}else{
			return -1
		}
	}

return final
}
func gcd(num1 int, num2 int )int{
	num1_array := []int{}
	num2_array := []int{}

	for i := 1 ; i < num1 ; i++{
		if num1 % i == 0{
			num1_array = append(num1_array, i)
		}
	}

	for i := 1 ; i < num2 ; i++{
		if num2 % i == 0{
			num2_array = append(num2_array, i)
		}
	}
	
	for i := len(num1_array)-1 ; i >=0 ; i--{
		for j := len(num2_array)-1 ; j >=0 ; j--{
			if num1_array[i] == num2_array[j]{
				return num1_array[i]
			}
		}
	}
	return 0
}
func gcd_chat(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}
func main() {
	// validation and convert number taking
	if len(os.Args) != 3 {
		return
	}
	pre_num_1 := os.Args[1]
	pre_num_2 := os.Args[2]

	if pre_num_1 == ""  || pre_num_2 == ""{
		return
	}
	num_1 := convert3(pre_num_1)
	num_2 := convert3(pre_num_2)
	if num_1 == -1  || num_2 == -1{
		return
	}

	fmt.Println(gcd(num_1,num_2))


}