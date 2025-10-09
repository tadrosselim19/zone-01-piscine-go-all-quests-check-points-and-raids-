package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	if len(os.Args) == 1{
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2{
		fmt.Println("Too many arguments")
		return
	}

	file,err:= os.Open(os.Args[1])
	if err != nil{
		return
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}