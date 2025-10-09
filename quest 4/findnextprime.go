package main

import "fmt"

func FindNextPrime(nb int) int {
	if nb <= 2{
		return 2
	}
	for i := 2 ; i*i <= nb ; i++{
		if nb%i == 0 {
			return FindNextPrime(nb-1)
		}
	}
	return nb
}

func aprimeFactors(n int) {
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			fmt.Print(i, " ")
			n /= i
		}
	}
	if n > 1 {
		fmt.Print(n)
	}
	fmt.Println()
}
func primeFactors(n int) []int{
	factsclice{
		
	}
}
func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(100))
	num := 225225
	fmt.Print("Prime factors of", num, ": ")
	primeFactors(num)
}