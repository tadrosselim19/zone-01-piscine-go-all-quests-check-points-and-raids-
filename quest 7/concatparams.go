package main

import (
	"fmt"
)

func ConcatParams(args []string) string {
	// names := os.Args[1:]
	final := ""
	for i := range args{
		final += args[i] + "\n"
	}
	return final
}
func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}