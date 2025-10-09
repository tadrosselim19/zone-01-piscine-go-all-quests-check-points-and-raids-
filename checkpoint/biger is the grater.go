package main

import "fmt"



func biggerIsGreater(w string) string {
	runes := []rune(w)
	n := len(runes)

	// Step 1: Find the pivot (rightmost character that is smaller than the next one)
	i := n - 2
	for i >= 0 && runes[i] >= runes[i+1] {
		i--
	}
	if i < 0 {
		return "no answer" // word is the highest permutation already
	}

	// Step 2: Find the smallest character to the right of pivot that is greater than runes[i]
	j := n - 1
	for runes[j] <= runes[i] {
		j--
	}

	// Step 3: Swap pivot with that character
	runes[i], runes[j] = runes[j], runes[i]

	// Step 4: Reverse the suffix (make it smallest possible)
	for left, right := i+1, n-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}

	return string(runes)
}

func main() {
	fmt.Println(biggerIsGreater("ab"))
	fmt.Println(biggerIsGreater("bb"))
	fmt.Println(biggerIsGreater("hefg"))
	fmt.Println(biggerIsGreater("dhck"))
	fmt.Println(biggerIsGreater("dkhc"))
	fmt.Println(biggerIsGreater("lmno"))
	fmt.Println(biggerIsGreater("dcba"))
	fmt.Println(biggerIsGreater("dcbb"))
	fmt.Println(biggerIsGreater("abcd"))
	fmt.Println(biggerIsGreater("fedcbabcd"))
}