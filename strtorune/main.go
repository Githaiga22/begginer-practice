package main

import (
	"fmt"
)

func main() {
	str := "Hello, 世界"
	for _, ch := range str {
		fmt.Printf("Rune: %c, Code Point: %U\n", ch, ch)
	}
}
