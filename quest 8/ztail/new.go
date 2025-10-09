package main

import (
	"fmt"
	"os"
)

func Atoi(s string) (int, bool) {
	result := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0, false
		}
		result = result*10 + int(ch-'0')
	}
	return result, true
}

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		fmt.Fprintln(os.Stderr, "Usage: go run . -c <bytes> <file1> [file2 ...]")
		os.Exit(1)
	}

	nBytes, ok := Atoi(os.Args[2])
	if !ok {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid byte count")
		os.Exit(1)
	}

	files := os.Args[3:]
	hasError := false
	multiple := len(files) > 1

	for i, name := range files {
		file, err := os.Open(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			hasError = true
			continue
		}

		info, err := file.Stat()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			file.Close()
			hasError = true
			continue
		}

		if multiple {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", name)
		}

		size := info.Size()
		start := size - int64(nBytes)
		if start < 0 {
			start = 0
		}

		_, err = file.Seek(start, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			file.Close()
			hasError = true
			continue
		}

		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)
			if n > 0 {
				os.Stdout.Write(buf[:n])
			}
			if err != nil {
				break
			}
		}

		file.Close()
	}

	if hasError {
		os.Exit(1)
	}
}
