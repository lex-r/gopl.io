package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io/ioutil"
	"os"
)

func main() {

	s512 := flag.Bool("sha512", false, "check sha512 checksum")
	s384 := flag.Bool("sha384", false, "check sha384 checksum")
	flag.Parse()

	var h hash.Hash

	switch {
	case *s512 == true:
		h = sha512.New()
	case *s384 == true:
		h = sha512.New384()
	default:
		h = sha256.New()
	}

	args := flag.Args()

	if len(args) == 0 {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Errorf("Cannot read data: %v\n", err)
			os.Exit(1)
		}
		checksum := sum(h, data)
		fmt.Printf("%x\n", checksum)
	} else {
		for _, fileName := range args {
			f, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Printf("%s: %s\n", fileName, err)
				continue
			}

			checksum := sum(h, f)
			fmt.Printf("%x - %s\n", checksum, fileName)

		}

	}
}

func sum(h hash.Hash, data []byte) []byte {
	h.Write(data)

	return h.Sum(nil)
}
