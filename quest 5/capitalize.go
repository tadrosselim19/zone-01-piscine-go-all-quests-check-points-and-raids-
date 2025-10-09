package main

import (
	"fmt"
	"unicode"
)
func aCapitalize(s string) string {
	srunes :=[]rune(s)
	result := ""
	
		if srunes[0] >= 'A' && srunes[0] <= 'Z'{
			result += string(srunes[0])
		}else if srunes[0] >= 'a' && srunes[0] <= 'z'{
			srunes[0]= srunes[0] - 32
			result += string(srunes[0])
		}
	for i := 1 ; i < len(s); i++{	
		if srunes[i] == ' '|| srunes[i] == '+'{
			if srunes[i+1] >= 'a' && srunes[i+1] <= 'z'{
				srunes[i+1]= srunes[i+1] - 32
				result += string(srunes[i+1])
			}else{
			result += string(srunes[i])
		    }
		}
	}
return result
}
func bCapitalize(s string) string {
	srunes := []rune(s)
	newWord := true

	for i, r := range srunes {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if newWord {
				srunes[i] = unicode.ToUpper(r)
			} else {
				srunes[i] = unicode.ToLower(r)
			}
			newWord = false
		} else {
			newWord = true
		}
	}
	return string(srunes)
}
func Capitalize(s string) string {
	srunes := []rune(s)
	isNewWord := true

	for i := 0; i < len(srunes); i++ {
		if (srunes[i] >= 'a' && srunes[i] <= 'z') || (srunes[i] >= 'A' && srunes[i] <= 'Z') || (srunes[i] >= '0' && srunes[i] <= '9') {
			if isNewWord && srunes[i] >= 'a' && srunes[i] <= 'z' {
				srunes[i] = srunes[i] - 32 // to uppercase
			} else if !isNewWord && srunes[i] >= 'A' && srunes[i] <= 'Z' {
				srunes[i] = srunes[i] + 32 // to lowercase
			}
			isNewWord = false
		} else {
			isNewWord = true
		}
	}

	return string(srunes)
}


func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}