package main

import "github.com/01-edu/z01"

func aPrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	var comb []int
	printCombNRecursive(n, 0, comb)
	z01.PrintRune('\n')
}

func printCombNRecursive(n, start int, comb []int) {
	if len(comb) == n {
		for _, d := range comb {
			z01.PrintRune(rune('0' + d))
		}
		if !isLastCombination(comb, n) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for i := start; i <= 9; i++ {
		if 10-i < n-len(comb) {
			break
		}
		printCombNRecursive(n, i+1, append(comb, i))
	}
}

func isLastCombination(comb []int, n int) bool {
	for i, d := range comb {
		if d != i+10-n {
			return false
		}
	}
	return true
}


func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var comb []int
	var generate func(pos, start int)

	generate = func(pos, start int) {
		if pos == n {
			for _, d := range comb {
				z01.PrintRune(rune('0' + d))
			}
			if comb[0] != 10-n {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			return
		}

		for i := start; i <= 9; i++ {
			if 10-i < n-pos {
				break
			}
			comb = append(comb, i)
			generate(pos+1, i+1)
			comb = comb[:len(comb)-1]
		}
	}

	generate(0, 0)
	z01.PrintRune('\n')
}

func main() {
	PrintCombN(1)
	println()
	PrintCombN(3)
	println()
	PrintCombN(9)
}