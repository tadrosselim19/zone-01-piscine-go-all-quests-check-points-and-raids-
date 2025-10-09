package main

import (
	"fmt"
	
)
func RetainFirstHalf(str string) string {
	if len(str) == 1{
		return str
	}
	str_rune := []rune(str)
	final := ""
	for i:= 0 ; i < len(str_rune)/2 ; i++{
		final += string(str_rune[i])
	}
	return final

}
func RetainFirstHalf(str string) string {
    runes := []rune(str)
    n := len(runes)
    if n == 0 {
        return ""
    }
    if n == 1 {
        return str
    }
    return string(runes[:n/2])
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
	fmt.Println(RetainFirstHalf("你好世界")) 

}