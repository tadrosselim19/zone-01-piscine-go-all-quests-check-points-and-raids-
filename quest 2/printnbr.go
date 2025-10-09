package main

import (

	"github.com/01-edu/z01"
)
func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	
	if n < 0{
		z01.PrintRune('-')
		var digits[]int
		
		for n < 0{
			digits = append(digits, n % 10)
			n = n / 10
		}
		for i := len(digits)-1 ; i >= 0 ; i--{
			z01.PrintRune(rune(-digits[i]+'0'))
		}
	}
	if n > 0{
		var digit[]int
		for n > 0{
			digit = append(digit, n % 10)
			n = n / 10
		}
		for  i := len(digit)-1 ; i >= 0 ; i--{
			z01.PrintRune(rune(digit[i]+'0'))
		}
	}
}
func cPrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		printDigits(-n)
		return
	}
	printDigits(n)
}

func printDigits(n int) {
	if n <= -10 || n >= 10 {
		printDigits(n / 10)
	}
	d := n % 10
	if d < 0 {
		d = -d
	}
	z01.PrintRune(rune('0' + d))
}
func aPrintNbr(n string) {
	nbyte := []rune(n)
	for i :=0 ; i < len(nbyte) ; i++{
		if (nbyte[i] > '9' || nbyte[i] < '0' ) && nbyte[i] != '-'{
			return 
		}else{
			z01.PrintRune(nbyte[i])
		}
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	PrintNbr(-9223372036854775808)
	z01.PrintRune('\n')
}