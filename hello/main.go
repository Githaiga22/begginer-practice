package main

import (
	"github.com/01-edu/z01"
)

func main() {
	message := "Hello World!"
	for _, char := range message {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
