package main

import (
	"fmt"
	"strconv"
)
func NotDecimal(dec string) string {
	if dec == "" || dec == "." ||dec == "-" {
		return "\n"
	}
	count_dot := 0
	for i:= 0 ; i < len(dec) ; i++ {
		if dec[i]=='.'{
		count_dot++
	}
	if count_dot > 1{
			return dec + "\n"
		}
	if (dec[i] > '9' || dec[i] <'0' )&& dec[i] != '.' && dec[i] != '-' {
			return dec + "\n"
		}

		
	}
	pre_final := ""
	for i:= 0 ; i < len(dec) ; i++ {
		if i==0 && dec[i]== '-'{
			pre_final += string(rune(dec[i]))
			continue
		}
		if dec[i] == '.'{
			count:= 0
			for  j := i+1 ; j < len(dec) ; j++{
			    if dec[j] == '0'{
				count++
				}
			}
			if count == len(dec)-i-1 && count != 0{
				break
			}else{
				continue
			}

		    }else{
			pre_final += string(rune(dec[i]))
		}
	}
	
	dec_final , err := strconv.Atoi(pre_final)
	if err != nil{
		return dec + "\n"
	}
	final := strconv.Itoa(dec_final)

	return  final + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
	fmt.Print(NotDecimal("25.00"))
	fmt.Print(NotDecimal("25.00999"))
	fmt.Print(NotDecimal("12.3.4"))
fmt.Print(NotDecimal("."))
fmt.Print(NotDecimal("-19.00"))
fmt.Print(NotDecimal("-0.0"))
fmt.Print(NotDecimal("00012.30"))
fmt.Print(NotDecimal(".25"))
fmt.Print(NotDecimal("25.."))
fmt.Print(NotDecimal("12-3.4"))
fmt.Print(NotDecimal("-"))
}