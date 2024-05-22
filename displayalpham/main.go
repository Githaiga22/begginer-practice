package main

import (
	"github.com/01-edu/z01"
)
func main() {
	for i := 'a'; i <= 'z'; i++ {
		if (i- 'a')% 2 == 0 {
			z01.PrintRune(i)
		} else {
			z01.PrintRune(i-'a' + 'A')
		}
	}
}