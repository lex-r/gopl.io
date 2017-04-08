package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for i := 1; i < len(os.Args); i++ {
		s += fmt.Sprintf("%d %s\n", i, os.Args[i])
	}
	fmt.Print(s)
}
