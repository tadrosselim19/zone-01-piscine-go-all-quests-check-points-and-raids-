package main

import (
	"fmt"
)
func is_pascal(s string)bool{
	for i:= 0 ; i < len(s); i++{
		if i == 0 && !(s[i] >= 'A' && s[i] <= 'Z'){
			return false
		}
		if!((s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z')){
			return false
		}
		if i+1 < len(s) && (s[i] >= 'A' && s[i] <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z'){
			return false
		}
	}
	return true
}
func SetSpace(s string) string {
	if s == "" {
    return ""
    }

	if is_pascal(s)==false{
		return "Error"
	}
	final:=""
	s_rune := []rune(s)
	for i:=0 ; i<len(s_rune);i++{
		if i !=0 && s_rune[i] >= 'A' && s_rune[i] <= 'Z'{
			final+=" "+ string(s_rune[i])
		}else{
			final+= string(s_rune[i])
		}
	}
	return final

}

func main() {
	fmt.Println(SetSpace("HelloWorld"))
	fmt.Println(SetSpace("HelloWorld12"))
	fmt.Println(SetSpace("Hello World"))
	fmt.Println(SetSpace(""))
	fmt.Println(SetSpace("LoremIpsumWord"))
	fmt.Println(SetSpace("HelloWORLD"))

}
