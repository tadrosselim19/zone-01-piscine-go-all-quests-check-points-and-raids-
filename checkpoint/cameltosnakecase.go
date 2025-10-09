package main

import (
	"fmt"
)
func validation_of_string(s string)bool{
	for i:= 0; i < len(s) ; i++{
		if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z'{
			return false
		}
		if !((s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z')){
			return false
		}
		if i !=len(s)-1 && ((s[i] >= 'A' && s[i] <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z')){
			return false
		}
	}
	return true
}
func CamelToSnakeCase(s string)string{
	validated_s := validation_of_string(s)
	if validated_s == false{
		return s
	}
	final := ""

	for i := 0 ; i < len(s) ; i++{
		if i !=0 && s[i] >= 'A' && s[i] <= 'Z'{
			final += "_"
			final += string(rune(s[i]))
		}else{
			final += string(rune(s[i]))
		}
	}
	return final
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
	fmt.Println(CamelToSnakeCase("CamelCAse"))
	
}

// $ go run .
// Hello_World
// hello_World
// camel_Case
// CAMELtoSnackCASE
// camel_To_Snake_Case
// hey2