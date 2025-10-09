package main

import "fmt"



func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	result := 1
	for i:= 1 ; i <= nb ; i++ {
		result = result * i
	}
	return result
}
func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	result := 1
	if nb > 1 {
		result = RecursiveFactorial((nb -1))* nb
	}
	return result
	
	
}
func aRecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
	fmt.Println(RecursiveFactorial(arg))

}

