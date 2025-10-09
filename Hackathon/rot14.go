package main

import (

	"github.com/01-edu/z01"
)
func Rot14(s string) string {
	result := ""
	srunes := []rune(s)
	for i := 0 ; i < len(srunes) ; i++{
		if (srunes[i] >= 'a' && srunes[i] <= 'z'){
			rotated := srunes[i] + 14
			if rotated > 'z' {
				rotated -= 26
			}
			result += string(rotated)
		}else if (srunes[i] >= 'A' && srunes[i] <= 'Z') {
			rotated := srunes[i] + 14
			if rotated > 'Z' {
				rotated -= 26
			}
			result += string(rotated)
		}else{
			result += string(srunes[i])
		}
		
	}
	return result 

}
func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}