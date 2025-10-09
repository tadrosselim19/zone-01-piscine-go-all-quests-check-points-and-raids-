package main

import (
    "fmt"
)
func ReverseMenuIndex(menu []string) []string {
	final := make([]string,len(menu))
	for i:= 0 ; i < len(final) ; i++{
		final[i] = menu[len(menu)-1-i]
	}
	return final
}

func main() {
        fmt.Println(ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
}
