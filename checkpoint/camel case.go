package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// Check if the string is not camelCase
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= '0' && c <= '9') || c == '_' || c == '-' {
			return s
		}
		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return s
		}
	}
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return s
	}

	result := ""
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			if i != 0 {
				result += "_"
			}
			result += string(c + 32) // convert to lowercase (ASCII trick)
		} else {
			result += string(c)
		}
	}

	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))       // Hello_World
	fmt.Println(CamelToSnakeCase("helloWorld"))       // hello_World
	fmt.Println(CamelToSnakeCase("camelCase"))        // camel_Case
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE")) // CAMELtoSnackCASE
	fmt.Println(CamelToSnakeCase("camelToSnakeCase")) // camel_To_Snake_Case
	fmt.Println(CamelToSnakeCase("hey2"))             // hey2
}
