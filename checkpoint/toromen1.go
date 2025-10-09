package main

import (
	"os"
	"fmt"
)

func convert1(num string)int{
	rune_num := []rune(num)
	final := 0	
	for i := 0 ;i < len(rune_num); i++{
		if rune_num[i] <= '9' && rune_num[i] >= '0'{
			final = final *10 + int(rune_num[i])-'0'
		}else{
			return 0
		}
	}

return final
}
func main(){
	if len(os.Args) != 2{
		fmt.Printf("ERROR: cannot convert to roman digit")
		return
	}
	pre_number := os.Args[1]
	post_number := convert1(pre_number)

	roman_array:=[]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman_num := []int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	if post_number == 0 || post_number > 4000{
		fmt.Printf("ERROR: cannot convert to roman digit")
		return
	}

	for post_number > 0{
		for i:=0 ; i < len(roman_array) ; i++{
			if post_number - roman_num[i] >= 0{
				fmt.Print(roman_array[i])
				post_number = post_number - roman_num[i]
				break
			}
		}
		
	}
	fmt.Println()



}