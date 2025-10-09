package main

import (
	"github.com/01-edu/z01"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1{
		for _, i := range "Hello\nHello"{
			z01.PrintRune(i)
		}
		os.Exit(1)
	}
	if len(os.Args) >= 2{
		for i := 1 ; i < len(os.Args) ;i++{
			file,err := os.Open(os.Args[i])
			if err != nil{
			   for _, i := range ("ERROR: open" + (os.Args[i]) + ": no such file or directory"){
			       z01.PrintRune(i) 
		        }
			}
			defer file.Close()
			io.Copy(os.Stdout,file)
			os.Exit(1)
		}
	}
}