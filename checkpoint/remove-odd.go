package main

import "fmt"
func RemoveOdd(s string) string{
	final := ""
	count := 0
	for i := 0 ; i < len(s) ; i++{
		count++
		if s[i] == ' '{
			final+=" "
			count= 0
			continue
		}
		if count % 2 != 0{
			final +=string(rune(s[i]))
		}
	}
    return final
}
func main(){
    fmt.Println(RemoveOdd("Hello World"))
    fmt.Println(RemoveOdd(" H"))
    fmt.Println(RemoveOdd("How are you?"))
	fmt.Println(RemoveOdd("Go  Lang"))

}