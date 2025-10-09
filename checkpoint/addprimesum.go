package main

import (
	"os"
	"github.com/01-edu/z01"
)
func convert2(num string)int{
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
func findprimesbelow(number int)int{
	calculation := 0
	if number <2 {
		return 0
	}
	for i := 2 ; i <= number ; i++{
		if isprime(i) == true{
			calculation += i
		}
	}
	return calculation 
}
func isprime(prime int)bool{
	for i:=2 ;i*i <= prime ; i++{
		if prime%i == 0{
			return false
		}
	}
	return true
}
func printrunesum(n int)[]int{
	nsclice := []int{}
	for n > 0{
		nsclice = append(nsclice, n%10)
		n = n / 10
	}
	return nsclice
}
func main(){
	if len(os.Args) != 2{
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	number_pre := os.Args[1]
	number_post := convert2(number_pre)
	
	if number_post <= 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	sum := findprimesbelow(number_post)
	revesed_sum :=printrunesum(sum)

	for i := len(revesed_sum)-1 ; i >= 0  ; i--{
		z01.PrintRune(rune(revesed_sum[i])+'0')
	}
	z01.PrintRune('\n')
}