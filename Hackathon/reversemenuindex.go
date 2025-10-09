package main

import (
        "fmt"
)
func ReverseMenuIndex(menu []string) []string {
	reverse := make([]string, len(menu))
	for i := 0 ; i < len(menu) ; i++{
		reverse[i]= menu[len(menu)-1-i]
	}
	return  reverse
}
func main() {
        fmt.Println(ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
}