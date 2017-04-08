package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	for _, fileName := range files {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		if hasDuplicateLines(f) {
			fmt.Println(fileName)
		}
		f.Close()
	}
}

func hasDuplicateLines(f *os.File) bool {
	lines := make(map[string]bool)
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if _, exists := lines[line]; exists {
			return true
		}
		lines[line] = true
	}
	// Примечание: игнорируем потенциальные
	// ошибки из input.Err()

	return false
}
