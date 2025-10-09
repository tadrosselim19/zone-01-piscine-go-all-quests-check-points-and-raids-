package main

import "fmt"

func AddFront(s string, slice []string) []string {
    new_string := []string{s}
	for i:=0 ; i < len(slice) ; i++{
		new_string = append(new_string, slice[i])
	}
	return new_string
}
func main() {
    fmt.Println(AddFront("Hello", []string{"world"}))
    fmt.Println(AddFront("Hello", []string{"world", "!"}))
    fmt.Println(AddFront("Hello", []string{}))
}