package main

import (
	"fmt"
)

func ShoppingSummaryCounter(str string) map[string]int {
	summarycount := make(map[string]int)
	word := ""
	for i := 0 ; i < len(str) ;i++{
		if str[i] == ' '{
			if word != ""{
				summarycount[word]++
			     word = ""
			}
		}else{
			word += string(str[i])
		}
		
	
	}

	if word != ""{
		summarycount[word]++
	}
	return summarycount

}

func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}
func AShoppingSummaryCounter(str string) map[string]int {
	res := make(map[string]int)
	word := ""
	inWord := false

	for _, ch := range str {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			if inWord {
				res[word]++
				word = ""
				inWord = false
			} else {
				res[""]++
			}

			continue
		}

		word += string(ch)
		inWord = true
	}

	if inWord {
		res[word]++
	} else {
		res[""]++
	}

	return res
}