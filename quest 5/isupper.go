package main

import (
	"fmt"
)

func IsUpper(s string) bool {
	srunes:=[]rune(s)
	for _, i:= range(srunes){
		if i < 'A' || i > 'Z'{
			return false
	}
	}
	return true
}
func IsLower(s string) bool {
	lrunes:=[]rune(s)
	for _, i:= range(lrunes){
		if i < 'a' || i > 'z'{
			return false
	}
	}
	return true

}

func IsAlpha(s string) bool {
	srunes :=[]rune(s)
	for _, i :=range(srunes){
		if !((i >= 'a' && i <= 'z' )|| (i >= 'A' && i <= 'Z' ) || (i >= '0' && i <= '9' )){
			return false
		}
	}
	return true

}
func aIsAlpha(s string) bool {
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			continue
		} else if i >= 'A' && i <= 'Z' {
			continue
		} else if i >= '0' && i <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
func IsNumeric(s string) bool {
	srunes :=[]rune(s)
	for _, i := range(srunes){
		if !(i >= '0' && i <= '9'){
			return false
		}
	}
	return true

}
func S_print(s string){
 	print(s)
 }
func IsPrintable(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}