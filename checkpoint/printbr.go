package main

import (
	"github.com/01-edu/z01"
)
func convert_unlimited(num int)([]int,bool){
	final := []int{}
	nigtive := false
	if num == 0 {
		final = append(final, 0)
		return final , false
	}
	if num < 0{
		nigtive = true
		if num % 10 != 0{
			last_number := -(num % 10) 
			final = append([]int{last_number},final...)
			num = num / 10
		}else{
			final =append([]int{0},final... )
		}
		num = num * -1
	}
	for num > 0{
		a := num % 10
		final = append([]int{a},final...)
		num = num / 10
	}
	return final , nigtive
}
func PrintNbr(n int) {
	number , negative:= convert_unlimited(n) 
	for i := 0 ; i < len(number) ; i++{
		if negative == true && i == 0{
			z01.PrintRune('-')
		}
		z01.PrintRune(rune(number[i]+'0'))
	}

}
func main() {
//	PrintNbr(-123)
//	PrintNbr(0)
//	PrintNbr(123)
//	PrintNbr(-9223372036854775808)
	PrintNbr(-9223372036854775808)
//	PrintNbr(9223372036854775807)

	
	z01.PrintRune('\n')
}