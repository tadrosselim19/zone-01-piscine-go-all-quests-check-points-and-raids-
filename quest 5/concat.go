package main

import (
	"fmt"
	
)

func aConcat(str1 string, str2 string) string {
	result:= str1
	result += str2
	return result

}
func Concat(str1 string, str2 string) string {
	return str1 + str2
}
func main() {
	fmt.Println(Concat("Hello!", " How are you?"))

}