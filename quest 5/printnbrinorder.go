package main

import (
	"strconv"
	"github.com/01-edu/z01"
)

func bPrintNbrInOrder(n int) {
	result := ""
	nstring := strconv.Itoa(n)
	r_nstring := []rune(nstring)
		for j := '0'; j <= '9'; j++ {
			for _, i := range r_nstring {
			if i == j {
				result += string(j)
			}
		}
		}
	
	print(result)

}

func aPrintNbrInOrder(n int) {
	result := ""
	nstring := strconv.Itoa(n)
	rnstring := []rune(nstring)

	for j := '0'; j <= '9'; j++ {
		for _, i := range rnstring {
			if i == j {
				result += string(j)
			}
		}
	}

	//for _, r := range result {
		//z01.PrintRune(r)
	//}
}
func PrintNbrInOrder(n int) {
	if n < 0 {
		return 
	}
	if n == 0{
		z01.PrintRune('0')
	}
	var digits [] int
	for n > 0{
		digits = append(digits, n%10)
		n = n/10
	}
	for i := 0 ; i < len(digits) ; i++{
		for j := i + 1 ; j < len(digits) ; j++{
			if digits[i] > digits[j] {
				digits[j] , digits[i] = digits[i], digits[j]
			}
		}
		
	}

	for _, a := range (digits){
		z01.PrintRune(rune(a) + '0')
	}


}


func main() {
	PrintNbrInOrder(800)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}