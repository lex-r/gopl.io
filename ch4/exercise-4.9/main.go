package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type wordFreq struct {
	word string
	freq int
}

type orderedWordFreq []wordFreq

func (o orderedWordFreq) Len() int {
	return len(o)
}

func (o orderedWordFreq) Less(i, j int) bool {
	return o[i].freq < o[j].freq
}

func (o orderedWordFreq) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Specify the filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	wf := wordfreq(file)
	owf := make(orderedWordFreq, 0)

	for word, freq := range wf {
		owf = append(owf, wordFreq{word, freq})
	}

	sort.Sort(owf)

	fmt.Println("word freq:")
	for _, w := range owf {
		fmt.Printf("%s: %d\n", w.word, w.freq)
	}
}

func wordfreq(input io.Reader) map[string]int {
	res := make(map[string]int)
	scanner := bufio.NewScanner(input)
	scanner.Split(scanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		res[word]++
	}
	return res
}

// copypasted from bufio.ScanWords
func scanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !issep(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if issep(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func issep(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r)
}
