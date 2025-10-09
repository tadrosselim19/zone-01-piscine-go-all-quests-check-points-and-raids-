package main

import (
	"github.com/01-edu/z01"
)
func Rot14(s string) string {
	final := ""
	for i := 0 ; i < len(s) ; i++{
		if (s[i] >= 'a' && s[i] <= 'z') {
			final+= string(rune(((s[i]-'a'+14)%26+'a')))
		}else if s[i] >= 'A' && s[i] <= 'Z'{
			final+= string(rune((s[i]-'A'+14)%26+'A'))
		}else{
			final+= string(rune(s[i]))
		}
	}
	return final
}
func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}