package main

import (
	"fmt"

)
func aMap(f func(int) bool, a []int) []bool {
	count := len(a)
	boolenvar := make([]bool,count) 
	for i:= 0 ; i < count ; i++{
		if a[i] == 1{
			boolenvar[i] = false
		}
		if a[i] == 2{
			boolenvar[i] = true
		}
		if a[i] > 2 && a[i]%2 !=  0{
			boolenvar[i] = true
		}else{
			boolenvar[i] = false
		}
	}
	return boolenvar
}
func Map(f func(int) bool, a []int) []bool {
	count := len(a)
	boolenvar := make([]bool,count) 
	for i:=0 ; i < count ; i++{
		boolenvar[i] = f(a[i])
		}
	return boolenvar
}
func IsPrime(nb int) bool {
	if nb <= 1{
			return false
	}
	for i := 2 ; i<nb  ; i++{
		if nb%i == 0 {
			return false 
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}