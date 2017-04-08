package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("Time 1: %v\n", time.Since(start1).Nanoseconds())

	start2 := time.Now()
	strings.Join(os.Args[1:], " ")
	fmt.Printf("Time 2: %v\n", time.Since(start2).Nanoseconds())
}
