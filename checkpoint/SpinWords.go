package main

import "fmt"


func SpinWords(str string) string {
  final := ""
  s_str := []string{}
  
  for i :=0 ; i<len(str);i++{
    if str[i] != ' ' && str[i] != '\t'{
      final += string(rune(str[i]))
    }else{
      if final != ""{
        s_str = append(s_str,final)
        final =""
      }
    }
  }
  if final != ""{
        s_str = append(s_str,final)
        final =""
  }
  for i:=0 ;i <len(s_str) ;i++{

    if len(s_str[i]) >= 5{
		r_str := []rune(s_str[i])
      for j := 0 ; j < len(s_str[i])/2 ;j++{
        (r_str[j]) ,  r_str[len(s_str[i])-j-1] = r_str[len(s_str[i])-j-1] ,r_str[j]
      }
	  s_str[i] = string(r_str)
    }
  }
  
  for i:=0 ; i <len(s_str);i++{
    if i != len(s_str)-1{
    final += s_str[i] + " "
    }else{
		final+= s_str[i]
	}
  }
  return final 
}// SpinWords
func main() {
	fmt.Println(SpinWords("Hey fellow warriors"))

	
}