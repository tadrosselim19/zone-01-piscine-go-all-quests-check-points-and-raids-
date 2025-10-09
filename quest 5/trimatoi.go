package main

import (
	"fmt"
	
)
func aTrimAtoi(s string) int {
	srunes := []rune(s)
	sign := 1
	result := ""
	for i := 0 ; i < len(srunes) ; i++{
		if srunes[i] == '-'{
			sign = -1
		}
		if srunes[i]>= '0' && srunes[i] <= '9'{
			result += string(srunes[i])
		}
	}
	if len(result) == 0{
		return 0
	}
	var result_convert int
	for _, j := range(result){
		result_convert = result_convert * 10 + int(j - '0')
	}
	return result_convert * sign

}

func bTrimAtoi(s string) int {
	sign := 1
	foundDigit := false
	result := 0

	for _, r := range s {
		if r == '-' && !foundDigit {
			sign = -1
		}
		if r >= '0' && r <= '9' {
			foundDigit = true
			result = result*10 + int(r-'0')
		}
	}
	return result * sign
}
func TrimAtoi(s string) int {
	nbyte := []rune(s)
	result :=0 
	negative := 1
	for i :=0 ; i < len(nbyte) ; i++{
		if (nbyte[i] > '9' || nbyte[i] < '0' ) && nbyte[i] != '-'{
			continue 
		}else if nbyte[i] == '-'{
			for j := 0 ; j < i ; j++{
				if nbyte[j] <= '9' && nbyte[j] >= '0'{
					negative = negative * -1
					break
				}
			}
			negative = negative * -1
			continue
		}else{
			result = result*10 + int(nbyte[i]-'0')
		}
	}
	return result * negative
}
func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}