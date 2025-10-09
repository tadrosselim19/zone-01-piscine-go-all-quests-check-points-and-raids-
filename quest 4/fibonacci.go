package main

import "fmt"

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}

	fibonacci := []int{0, 1}
	for i := 2; i < index+1; i++ {
		fibonacci = append(fibonacci, fibonacci[i-2]+fibonacci[i-1])
	}
	return fibonacci[index]
}

func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}

func bFibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return bFibonacci(index-1) + bFibonacci(index-2)
}