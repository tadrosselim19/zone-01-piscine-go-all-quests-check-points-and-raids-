package main

import (
	"fmt"
	"io"
	"os"
)

func amain() {
	args := os.Args[1:]

	// If no arguments are given, read from stdin and write to stdout
	if len(args) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// If arguments are given, try to read and print each one
	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, file)
		file.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}
	}
}
package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			os.Stderr.Write([]byte("ERROR: " + err.Error() + "\n"))
			os.Exit(1)
		}
		return
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			os.Stderr.Write([]byte("ERROR: " + err.Error() + "\n"))
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, file)
		file.Close()

		if err != nil {
			os.Stderr.Write([]byte("ERROR: " + err.Error() + "\n"))
			os.Exit(1)
		}
	}
}
package main

import (
	"github.com/01-edu/z01"
	"os"
)

// Print string using z01
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) == 1 {
		// No arguments: read from stdin and echo input
		buf := make([]byte, 1)
		for {
			_, err := os.Stdin.Read(buf)
			if err != nil {
				break
			}
			z01.PrintRune(rune(buf[0]))
		}
		return
	}

	// Read and print each file passed as an argument
	for i := 1; i < len(os.Args); i++ {
		file, err := os.Open(os.Args[i])
		if err != nil {
			// Print error using z01 only
			printStr("ERROR: open " + os.Args[i] + ": no such file or directory\n")
			os.Exit(1)
		}

		// Read file content byte by byte and print with z01
		buf := make([]byte, 1)
		for {
			_, err := file.Read(buf)
			if err != nil {
				break
			}
			z01.PrintRune(rune(buf[0]))
		}
		file.Close()
	}
}

