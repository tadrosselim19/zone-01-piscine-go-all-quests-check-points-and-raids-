package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrStr := os.Args[1]
	targetStr := os.Args[2]

	// Parse array input like "[1, 2, 3, 4, 5]"
	if !strings.HasPrefix(arrStr, "[") || !strings.HasSuffix(arrStr, "]") {
		fmt.Println("Invalid input.")
		return
	}

	arrStr = strings.TrimPrefix(arrStr, "[")
	arrStr = strings.TrimSuffix(arrStr, "]")
	strNums := strings.Split(arrStr, ",")

	var arr []int
	for _, s := range strNums {
		numStr := strings.TrimSpace(s)
		if numStr == "" {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Invalid input.")
			return
		}
		if num < -1000 || num > 1000 {
			fmt.Printf("Invalid number: %d\n", num)
			return
		}
		arr = append(arr, num)
	}

	target, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	pairs := findPairs(arr, target)
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
	}
}

func findPairs1(arr []int, target int) [][2]int {
	var result [][2]int
	n := len(arr)
	used := make([]bool, n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !used[i] && !used[j] && arr[i]+arr[j] == target {
				result = append(result, [2]int{i, j})
				used[i], used[j] = true, true
				break // only one usage per element
			}
		}
	}
	retur result
}
