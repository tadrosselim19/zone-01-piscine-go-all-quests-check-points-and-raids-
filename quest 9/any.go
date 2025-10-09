package main

import (
	"fmt"
)
func Any(f func(string) bool, a []string) bool {
	for i := 0 ; i < len(a) ; i++{
		if f(a[i]) == true{
			return true
		}
	}
	return false

}
func IsNumeric(s string)bool{
	for _, i := range s{
		if i >= '0' && i <='9'{
			return true
		}
	}
	return false
}
func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}