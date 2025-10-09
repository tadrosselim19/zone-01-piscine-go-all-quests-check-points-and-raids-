package main

import (
	"fmt"
)
func LoafOfBread(str string) string {
	final := ""
	r_str := []rune(str)
	len := len(r_str)
	if len < 5{
		return "Invalid Output\n"
	}
	count := 0

	for i := 0 ; i < len ;i++{
		if count < 5 && r_str[i] == ' '{
			continue
		}
		if count == 5 {
			final += " "
			count = 0
		}else{
			final += string(r_str[i])
			count++
		}
	}
	return final+"\n"

}
func main() {
	fmt.Print(LoafOfBread("deliciousbread"))
	fmt.Print(LoafOfBread("This is a loaf of bread"))
	fmt.Print(LoafOfBread("loaf"))
}