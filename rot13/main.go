package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// checking no of arguments passed is only 2, prog name and input string
	if len(os.Args) != 2 {
		return
	}
	// initialize our input which is the command line argument 1
	input := os.Args[1]

	// range over our input characters
	for _, char := range input {
		// handling if incase the input char is uppercase
		if char >= 'A' && char <= 'Z' {
			// apply rot13 formulae and print uppercase letters
			z01.PrintRune((char-'A'+13)%26 + 'A')
			// check if lowercase as well
		} else if char >= 'a' && char <= 'z' {
			// apply rot13 formulae and print lowercase letters
			z01.PrintRune((char-'a'+13)%26 + 'a')
		} else {
			// else also ptint non alphabetical characters
			z01.PrintRune(char)
		}
	}
	// print on a newline
	z01.PrintRune('\n')
}
