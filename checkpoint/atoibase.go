package main

import (
	"fmt"
	
)
func validated_s(s string ,base string)bool{
	flag := false
	for i := 0 ; i < len(s) ; i++{
		for j := 0 ; j < len(base);j++{
			if s[i] == base[j]{
				flag = true
				break
			}
		}
		if flag == false{
			return  false
		}	
	}
	return flag
}
func mathpower(a,b int)int{
	result := 1
	for i := 0 ;i < b ;i++{
		result = result *a
	}
	return result
}
func AtoiBase(s string ,base string) int {
	// validation of base and s string
	if len(base) <2 || validated_s(s,base)== false{
		return 0
	}
	for i := 0 ; i < len(base);i++{
		if base[i] == '-' || base[i] == '+'{
			return 0
		}
	}
	for i := 0; i < len(base); i++ {
	   for j := i+1; j < len(base); j++ {
		if base[i] == base[j] {
			return 0
		}
	  }
    }

	// main function stared
	final := 0
	for i := 0 ; i < len(s) ; i++{
		for j := 0 ; j < len(base);j++{
			if s[i] == base[j]{
				final += j * mathpower((len(base)),(len(s)-i-1))
			}
		}
	}
	return final
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}