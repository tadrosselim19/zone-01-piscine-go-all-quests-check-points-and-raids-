package main

import (
	"fmt"
)
func Atoi(s string) int {
	if len(s) == 0 {
    return 0
    }

	final := 0 
	negative := 1
	for i:= 0 ; i < len(s) ; i++{
		// checking +  and - sign repetation
		if i == 0 && (s[0] == '+' || s[0] == '-') && (s[1] > '9' || s[1] <'0'){
			return 0
		}
		if i == 0 && s[0] == '+' {
            continue
        }
        if i == 0 && s[0] == '-' {
            negative = -1
            continue
        }

		if s[i] >= '0' && s[i] <= '9'{
			final = final *10 + int(s[i]-'0')
		}else{
			return 0
		}


	}
	return negative* final

}
func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("---1234"))

}
