package main

import (
	"fmt"
)
func ShoppingSummaryCounter(str string) map[string]int {
	final := map[string]int{}
	word := ""
	for i:= 0 ; i < len(str) ; i++{
		if str[i] != ' ' && str[i] != '\t'{
			word+= string(rune(str[i]))
		}else{
			if word != "" {
				final[word]+=1
				word = ""
			}
		}
	}
	if word != "" {
		final[word]+=1
	}
	return final
}

func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}
