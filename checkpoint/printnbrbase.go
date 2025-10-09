package main

import (
	"github.com/01-edu/z01"
)
func convert_base_unlimited(num int , length_base int)([]int, bool){
	final := []int{}
	negative := false
	// transation negtive number to positive to fit the index
	if num * -1 > 0{
		negative = true
		num = num * -1
	}
	// find the lagest math power
	x := 1
	for i := 1 ; i < num ; i = i* length_base{
		x = i
	}
	// looping for find the index
	for x >=1 {
		index := num / x
		if num == x || index >= length_base{
			index = length_base-1
	    }
		final = append(final, index)
		num = num % x
		x = x / length_base
	}
	return final , negative
}
func PrintNbrBase1(nbr int, base string) {
	length_base := len(base)
	num := nbr
	// validation of base
	// 1- length validation
	if length_base < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')

		return
	}
	// 2- chacters must be unique and not contain - or +
	for i := 0 ; i < length_base ; i ++{
			if base[i] == '+' || base[i] == '-' {
            z01.PrintRune('N')
            z01.PrintRune('V')
            return
        }
		for j := 0 ; j < length_base ; j ++{
		
			if base[i] == base[j] && i != j{
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
	   }
	}
	// Main function of printing
	number , negative := convert_base_unlimited(num,length_base)
	if negative == true{
		z01.PrintRune('-')
	}
	for i:= 0 ; i < len(number) ; i ++{
		z01.PrintRune(rune(base[number[i]]))
	}

}


func main() {
	PrintNbrBase1(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase1(0, "0123456789")
    z01.PrintRune('\n')
	PrintNbrBase1(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase1(-2147483648, "01")
	z01.PrintRune('\n')
	PrintNbrBase1(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase1(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase1(125, "aa")
	z01.PrintRune('\n')
	PrintNbrBase1(5, "a")
	z01.PrintRune('\n')
	PrintNbrBase1(-9223372036854775808, "01")
    z01.PrintRune('\n')
}

