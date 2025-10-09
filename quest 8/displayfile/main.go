package main

import (
	"fmt"
	"os"
	"io"
)

func main(){
	file ,err := os.Open("quest8.txt")
	if err != nil{
		fmt.Printf("File name missing")
	}
	arg := os.Args[1:]
	file.Read(arg[]byte,)package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}

	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println(err)
	}
}

}