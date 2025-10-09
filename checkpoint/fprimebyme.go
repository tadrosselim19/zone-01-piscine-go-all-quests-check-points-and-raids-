package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)
func is_prime(num int)bool{
	if num == 2{
		return true
	}
	if num <2{
		return false
	}
	for i := 2 ; i <num ; i++{
		if num%i ==0 {
			return false
		}
	}
	return true
}
func fprime(target int)([]int){
	final := []int{}
	i := 2
	for target != 1 {
			if target % i == 0{
				final = append(final, i)
				target = target / i 
			}else{
				i++
			    for is_prime(i)==false{
					i++
			    }
				} 
			} 
	return final
}
func main() {
	if len(os.Args) != 2{
		return 
	}
	target , err := strconv.Atoi(os.Args[1])
	if err != nil{
		return
	}
	if target < 2{
		return
	}
	// printing the result 
	factors := fprime(target)
	final := ""
	// cnvert to fit useing of z01.printrune
	for i := 0 ; i < len(factors) ; i++{
		if i != len(factors)-1{
			final += strconv.Itoa(factors[i]) + "*"
		}else{
		    final += strconv.Itoa(factors[i])
		}
	}

	for _, ch := range(final) {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
	

}