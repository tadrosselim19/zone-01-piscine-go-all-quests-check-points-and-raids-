package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	// split input by whitespace
	tokens := strings.Fields(os.Args[1])
	if len(tokens) == 0 {
		fmt.Println("Error")
		return
	}

	stack := []int{}

	for _, tok := range tokens {
		// try number
		if num, err := strconv.Atoi(tok); err == nil {
			stack = append(stack, num) // push
			continue
		}

		// must be operator
		if len(stack) < 2 {
			fmt.Println("Error")
			return
		}
		b := stack[len(stack)-1]
		a := stack[len(stack)-2]
		stack = stack[:len(stack)-2] // pop 2

		if tok == "+" {
			stack = append(stack, a+b)
		} else if tok == "-" {
			stack = append(stack, a-b)
		} else if tok == "*" {
			stack = append(stack, a*b)
		} else if tok == "/" {
			if b == 0 {
				fmt.Println("Error")
				return
			}
			stack = append(stack, a/b)
		} else if tok == "%" {
			if b == 0 {
				fmt.Println("Error")
				return
			}
			stack = append(stack, a%b)
		} else {
			fmt.Println("Error")
			return
		}
	}

	// after all tokens processed, stack must have exactly 1 value
	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}

	fmt.Println(stack[0])
}
