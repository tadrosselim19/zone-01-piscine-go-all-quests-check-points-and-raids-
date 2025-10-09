package main

import (
	"fmt"
)
func aiSplit(s, sep string) []string {
	newslice :=[]string{}
	word := ""
	for i := 0 ; i < len(s) ;{
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep{
			newslice = append(newslice, word)
			word = ""
			i += len(sep)
		}else{
			word += string(s[i])
			i++
		}
	}
	newslice = append(newslice, word)
	return newslice
}

func acSplit(s, sep string) []string {
	result := []string{}
	sepLen := len(sep)
	start := 0

	for i := 0; i <= len(s)-sepLen; {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			i += sepLen
			start = i
		} else {
			i++
		}
	}
	result = append(result, s[start:])
	return result
}

func mSplit(s, sep string) []string {
	// Step 1: Count how many times sep appears
	count := 1
	for i := 0; i+len(sep) <= len(s); {
		if s[i:i+len(sep)] == sep {
			count++
			i += len(sep)
		} else {
			i++
		}
	}

	// Step 2: Create slice with exact size
	result := make([]string, count)
	word := ""
	index := 0

	// Step 3: Fill the slice using indexing
	for i := 0; i < len(s); {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result[index] = word
			index++
			word = ""
			i += len(sep)
		} else {
			word += string(s[i])
			i++
		}
	}
	result[index] = word // add the last word

	return result
}

func Split(s, sep string) []string {
	count := 0
	word := ""
	
	for i := 0 ;i+len(sep) <= len(s) ;{
		if s[i:i+len(sep)] == sep{
			count++ 
			i += len(sep)
		}
		i++
	}
	all := make([]string,count+1)
	index := 0
	for i:=0 ; i < len(s) ;{
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep{
			all[index] = word
			index++
			word = ""
			i += len(sep)
		}
		word += string(s[i])
		i++

	}
	all[index]=word
	return all
}
func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}