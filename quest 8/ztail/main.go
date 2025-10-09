package main

import (
	"fmt"
	"os"
	"reflect"
)
func atoi(s string)(int , bool){
	result := 0
	for _, ch := range s{
		if ch > '9' || ch < '0'{
			return 0 , false
		}
		result = result*10 + int(ch - '0')
	}
	return result , true
}
func ExtractAllNumbers(s string) []int{
	allnumber := []int{}
	result := 0
	found := false

	for _, ch := range s{
		if ch <= '9' && ch >= '0'{
			 found = true
			 result = result*10 + int(ch - '0')
		}else if found{
         allnumber = append(allnumber, result)
		result = 0
		found = false
		}
	}
	if found{
		allnumber = append(allnumber, result)
	}
	return allnumber
}

func main(){
	if len(os.Args) < 4 || os.Args[1] == "-c"{
		fmt.Fprintln(os.Stderr,"Usage: go run . -c <bytes> <file1> [file2 ...]")
	}
	nbyte , ok := atoi(os.Args[2])
	if !ok {
		fmt.Fprintf(os.Stderr,"ERROR: Invalid byte count")
		os.Exit(1)
	}
	files := os.Args[3:]
	hasError := false
	multiple := len(files) > 1
	

}
