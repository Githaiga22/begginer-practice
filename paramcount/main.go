package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	argcount := len(os.Args) - 1

	strcount := strconv.Itoa(argcount)

	for _, char := range strcount {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
