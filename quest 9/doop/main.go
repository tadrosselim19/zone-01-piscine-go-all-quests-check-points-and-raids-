package main

import (
	"os"
)

func main(){
	if len(os.Args) != 4 {
		return
	}
	check1 := []byte(os.Args[1])
	    value1 := 0
		for j :=0 ; j < len(check1) ; j++{
			if check1[j] > '9' || check1[j] < '0'{
				return
			}
			value1 = value1*10 + int(check1[j] - '0')

		}
	if value1 < -9223372036854775808 || value1 > 9223372036854775807{
		return
	}

    check2 := []byte(os.Args[3])
	    value2 := 0
		for j :=0 ; j < len(check2) ; j++{
			
			if check2[j] > '9' || check2[j] < '0'{
				return
			}
			value2 = value2*10 + int(check2[j] - '0')
		}
	
	if value2 < -9223372036854775808 || value2 > 9223372036854775807{
		return
	}
	
	operatorcheck := os.Args[2]
	   if len(operatorcheck) != 1 {
		return
	   }
	   switch operatorcheck {
	   case "+":
		println(value1 + value2)
	   case "-":
		println(value1 - value2)
	   case "*":
		println(value1 * value2)
	   case "/":
		if  value2 == 0{
			println("No division by 0")
			return
		}
		println(value1 / value2)
	   case "%":
		if  value2 == 0{
			println("No division by 0")
			return
		}
		println(value1 % value2)
	   default:
          return
       }
}


