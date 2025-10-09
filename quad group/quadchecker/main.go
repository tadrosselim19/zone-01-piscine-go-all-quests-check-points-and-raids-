package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	arr, _ := ioutil.ReadAll(os.Stdin)
	if string(arr) == "erre" {
		fmt.Println("erre")
		return
	}

	arr1 := string(arr)
	if arr1[len(arr1)-1] != '\n' {
		arr1 += "\n"
	}
	var y, x int
	for i := 0; arr1[i] != '\n'; i++ {
		x++
	}
	for _, r := range arr1 {
		if r == '\n' {
			y++
		}
	}
	if arr1[0] == 'o' {
		cmd := exec.Command("./quadA", string(rune(x+'0')), string(rune(y+'0')))
		d, _ := cmd.Output()
		if quadcheck(arr1, string(d)) {
			fmt.Println("[quadA] [" + string(rune(x+'0')) + "] [" + string(rune(y+'0')) + "]")
		} else {
			fmt.Println("Not a quad function")
		}
	} else if arr1[0] == '/' {
		cmd := exec.Command("./quadB", string(rune(x+'0')), string(rune(y+'0')))
		d, _ := cmd.Output()
		if quadcheck(arr1, string(d)) {
			fmt.Println("[quadB] [" + string(rune(x+'0')) + "] [" + string(rune(y+'0')) + "]")
		} else {
			fmt.Println("Not a quad function")
		}
	} else if arr1[0] == 'A' {
		var t []string
		tr := []string{"quadC", "quadD", "quadE"}
		for i := 0; i < 3; i++ {
			cmd := exec.Command("./"+tr[i], string(rune(x+'0')), string(rune(y+'0')))
			d, _ := cmd.Output()
			if quadcheck(arr1, string(d)) {
				t = append(t, tr[i])
			}
		}
		if len(t) == 0 {
			fmt.Println("Not a quad function")
		} else {
			for i := 0; i < len(t); i++ {
				fmt.Print("["+t[i]+"] ["+string(rune(x+'0'))+"]", "["+string(rune(y+'0'))+"]")
				if i != len(t)-1 {
					fmt.Print(" || ")
				}
			}
			fmt.Println()
		}

	} else {
		fmt.Println("Not a quad function")
	}
}

func quadcheck(s string, r string) bool {
	if len(s) != len(r) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != r[i] {
			return false
		}
	}
	return true
}