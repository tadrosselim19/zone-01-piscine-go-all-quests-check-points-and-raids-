package main

import (
	"os"

	"github.com/01-edu/z01"
)
func convert6(num string)(int,bool){
	rune_num := []rune(num)
	final := 0
	flag := 1
	for i := 0 ;i < len(rune_num); i++{
		if rune_num[0] == '-'{
			flag *= -1
			continue
		}
		if rune_num[i] <= '9' && rune_num[i] >= '0'{
			final = final *10 + int(rune_num[i])-'0'
		}else{
			return -1 ,false
		}
	}
return final * flag ,true
}
func find_index(num int)[]int{
	final := []int{}
	index := 0
	x := 1
	for i := 1 ; num > i ; i*=16{
		x = i
	}
	for x >= 1{
		index = num/x
		final= append(final, index)
		num = num % x
		x = x/16
	}
	return final
}
func main() {
	if len(os.Args) != 2 {
		return
	}
	number ,err := convert6(os.Args[1])
	if err == false{
		os.Stdout.WriteString("ERROR\n")
		return
	}
	if number == 0 {
    z01.PrintRune('0')
    z01.PrintRune('\n')
    return
    }

	hex_base := "0123456789abcdef"
	test := find_index(number)

	for i := 0 ; i < len(test) ; i++{
		z01.PrintRune(rune(hex_base[test[i]]))
	}
	z01.PrintRune('\n')
	//hex_refrance := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
}