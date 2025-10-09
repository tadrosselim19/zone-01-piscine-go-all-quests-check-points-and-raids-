package main

import (

	"os"

	"github.com/01-edu/z01"
)
func is_repeted(str string, ch rune)bool{
	str_rune := []rune(str)
	for i:=0 ; i < len(str_rune) ; i++{
		if str_rune[i] == ch {
			return true
		}
	}
	return false
}
func main(){
	if len(os.Args) !=3 {
		z01.PrintRune('\n')
		return
	}
// join the argument
	pre_final := os.Args[1] + os.Args[2]
	final := ""
// the union function 
	pre_final_rune := []rune(pre_final)
	for i := 0 ; i < len(pre_final_rune) ; i++{
		if final == ""{
			final += string(pre_final_rune[0])
		}
		if is_repeted(final , pre_final_rune[i]){
			continue
		}else{
			final += string(pre_final_rune[i])
		}
	}
// printing using z01.printrune
for i:= 0 ; i < len(final) ; i++{
	z01.PrintRune(rune(final[i]))
}
z01.PrintRune('\n')
}