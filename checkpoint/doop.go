package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	operator_list := "/%*+-"
	fr_value := os.Args[1]
	operator := os.Args[2]
	se_value := os.Args[3]

// validation of argument 
	  // 1- validation of values
	fr_num , err1 := strconv.Atoi(fr_value)
	se_num , err2 := strconv.Atoi(se_value)
	if err1 != nil{
		return
	}
	if err2 != nil{
		return
	}
	if fr_num < -9223372036854775808 || fr_num > 9223372036854775807{
		return
	}
	if se_num < -9223372036854775808 || se_num > 9223372036854775807{
		return
	}
	  // 2- validation of operator
	if len(operator) != 1{
		return
	}
	operator_valid := false
	for i:= 0 ; i <len(operator_list) ;i++{
		if operator[0] == operator_list[i]{
			operator_valid=true
		}
	}
	if operator_valid == false{
		return
	}

// opperation 
   if operator == "+"{
	if fr_num >0 && se_num > 0 && fr_num + se_num < 0{
		return
	}else if fr_num < 0 && se_num < 0 && fr_num + se_num > 0{
		return
	}
	  fmt.Println(fr_num + se_num)

   }else if operator == "-" {
	if fr_num <0 && se_num < 0 && fr_num - se_num > 0{
	  return
	}
	  fmt.Println(fr_num - se_num)
   }else if operator == "*"{
        if fr_num > 0 && se_num > 0 && fr_num * se_num < 0{
		   return
	    }else if fr_num < 0 && se_num < 0 && fr_num * se_num < 0{
		   return
	    }else if fr_num >0 && se_num < 0 && fr_num * se_num > 0{
	     return
	    }else if fr_num <0 && se_num > 0 && fr_num * se_num > 0{
	     return
	    }
	  fmt.Println(fr_num * se_num)
   }else if operator == "/" && se_num != 0{
	  fmt.Println(fr_num / se_num)
   }else if operator == "/" && se_num == 0{
	  fmt.Println("No division by 0")
   }else if operator == "%" && se_num != 0{
	  fmt.Println(fr_num % se_num)
   }else if operator == "%" && se_num == 0{
	  fmt.Println("No modulo by 0")
   }





	


}