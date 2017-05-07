package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * KB
	GB = KB * MB
	TB = KB * GB
	PB = KB * TB
	EB = KB * PB
	ZB = KB * EB
	YB = KB * ZB
)

func main() {
	fmt.Printf("KB: %d\nMB: %d\nGB: %d\n", KB, MB, GB)
}
