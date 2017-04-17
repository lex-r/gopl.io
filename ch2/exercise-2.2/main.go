package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error parsing number: %v\n", err)
				os.Exit(1)
			}

			convert(num)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			arg := scanner.Text()
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error parsing number: %v\n", err)
				os.Exit(1)
			}
			convert(num)
		}
	}
}

func convert(number float64) {
	fahr := Fahrenheit(number)
	cels := Celsius(number)
	foot := Foot(number)
	metr := Metre(number)
	pound := Pound(number)
	kg := Kilogram(number)

	fmt.Printf("%s = %s, %s = %s\n",
		fahr, FToC(fahr), cels, CToF(cels))
	fmt.Printf("%s = %s, %s = %s\n",
		foot, FToM(foot), metr, MToF(metr))
	fmt.Printf("%s = %s, %s = %s\n",
		pound, PToK(pound), kg, KToP(kg))
}
