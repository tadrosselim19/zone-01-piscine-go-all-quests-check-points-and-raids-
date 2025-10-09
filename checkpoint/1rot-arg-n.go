package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3{
		fmt.Println()
		return
	}
	num , err:= strconv.Atoi(os.Args[1])
	if err != nil || num < 0 {
		fmt.Println("Error")
		return 
	}
	target := os.Args[2:]
	stared := num%len(os.Args[2:])
	if stared== 0 || num == 0{
		for i := 0 ; i < len(target) ;i++{
			if i != len(target)-1 {
				fmt.Print(target[i]+ " ")
			}else{
				fmt.Print(target[i])
			}
		
		}
		return
	}
	end := len(os.Args[2:])
	for i := stared ; i < end  ;{
		if i != len(target)-1 && stared != 0{
		   fmt.Print(target[i]+ " ")
		   i++
		}else if i == len(target)-1{
		   fmt.Print(target[i])
		   i = 0
		   end = stared
		   fmt.Print(" ")
		}
	}

}



package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println()
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num < 0 {
		fmt.Println("Error")
		return
	}

	args := os.Args[2:]
	shift := num % len(args)

	rotated := append(args[shift:], args[:shift]...)
	fmt.Println(strings.Join(rotated, " "))
}
