package main

func ForEach(f func(int), a []int) {
	for i := 0 ; i < len(a) ; i++ {
		f(a[i])
	}

}
func PrintNbr(nub int){
	print(nub)
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}
package piscine

import "github.com/01-edu/z01" // or another allowed method

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	printDigits(n)
}

func printDigits(n int) {
	if n == 0 {
		return
	}
	printDigits(n / 10)
	z01.PrintRune(rune(n%10 + '0'))
}
