package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("Can't open file")
	}
	return strings.TrimRight(string(content), "\n")
}

func main() {
	fmt.Println(day11())
}
