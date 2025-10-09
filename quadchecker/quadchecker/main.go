package main

import (
	"fmt"
	"io"
	"os"
)

func check(input string, width, height int, style []byte) bool {
	row, col := 0, 0

	for i := 0; i < len(input); i++ {
		ch := input[i]

		if ch == '\n' {
			if col != width {
				return false
			}
			row++
			col = 0
			continue
		}

		if row >= height || col >= width {
			return false
		}

		var expected byte
		top := row == 0
		bottom := row == height-1
		left := col == 0
		right := col == width-1

		switch {
		case top && left:
			expected = style[0]
		case top && right:
			expected = style[2]
		case bottom && left:
			expected = style[6]
		case bottom && right:
			expected = style[8]
		case top || bottom:
			expected = style[1]
		case left || right:
			expected = style[3]
		default:
			expected = style[4]
		}

		if ch != expected {
			return false
		}
		col++
	}

	return row == height
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil || len(data) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	input := string(data)

	// Parse width and height
	width, height := 0, 0
	lineLen := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			if width == 0 {
				width = lineLen
			} else if lineLen != width {
				fmt.Println("Not a quad function")
				return
			}
			height++
			lineLen = 0
		} else {
			lineLen++
		}
	}

	if width == 0 || height == 0 || lineLen != 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Define quad styles
	styles := []struct {
		name string
		set  []byte
	}{
		{"quadA", []byte{'o', '-', 'o', '|', ' ', '|', 'o', '-', 'o'}},
		{"quadB", []byte{'/', '*', '\\', '*', ' ', '*', '\\', '*', '/'}},
		{"quadC", []byte{'A', 'B', 'A', 'B', ' ', 'B', 'C', 'B', 'C'}},
		{"quadD", []byte{'A', 'B', 'C', 'B', ' ', 'B', 'A', 'B', 'C'}},
		{"quadE", []byte{'A', 'B', 'C', 'B', ' ', 'B', 'C', 'B', 'A'}},
	}

	matched := 0
	for _, s := range styles {
		if check(input, width, height, s.set) {
			if matched > 0 {
				fmt.Print(" || ")
			}
			fmt.Printf("[%s] [%d] [%d]", s.name, width, height)
			matched++
		}
	}

	if matched == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println()
	}
}
