package main

import (
	"fmt"
	
)

func SplitWhiteSpaces(s string) []string {
	newsclics := []string{}
	newword := ""
	for _, i := range s{
		if i != ' ' && i != '\t' && i != '\n' {
		   word := ""
		   word += string(i)
		   newword += word
		}
		if i == ' ' || i == '\t' || i == '\n'{
		   newsclics =append(newsclics, newword)
           word := ""
		   newword = word
		}
	}
	if newword != "" {
		newsclics =append(newsclics, newword)
	}
	return newsclics
}
func cSplitWhiteSpaces(s string) []string {
	newsclics := []string{}
	newword := ""

	for _, i := range s {
		if i != ' ' && i != '\t' && i != '\n' {
			word := ""
			word += string(i)
			newword += word
		} else {
			if newword != "" {
				newsclics = append(newsclics, newword)
				newword = ""
			}
		}
	}
	if newword != "" {
		newsclics = append(newsclics, newword)
	}
	return newsclics
}


func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
