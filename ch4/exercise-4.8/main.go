package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		switch {
		case unicode.IsControl(r):
			counts["controls"]++
		case unicode.IsDigit(r):
			counts["digits"]++
		case unicode.IsLetter(r):
			counts["letters"]++
		case unicode.IsNumber(r):
			counts["numbers"]++
		case unicode.IsPunct(r):
			counts["puncts"]++
		case unicode.IsSpace(r):
			counts["spaces"]++
		default:
			counts["unknowns"]++
		}
	}

	fmt.Print("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
