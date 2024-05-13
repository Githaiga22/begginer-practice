package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() { // function that displays the first parameter
	// check if there is atleast one command line argument
	if len(os.Args) > 1 {
		// iterate over the first characters of the first argument
		for _, char := range os.Args[1] {
			// print the first character
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
