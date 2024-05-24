package main

import "github.com/01-edu/z01"

// ReduceInt applies the function f to each element in the slice a and prints the result
func ReduceInt(a []int, f func(int, int) int) {
	// Initialize the accumulator with the first element of the slice
	acc := a[0]

	// Iterate over the remaining elements in the slice
	for _, num := range a[1:] {
		// Apply the function f to the accumulator and the current element
		acc = f(acc, num)
	}

	// Print the final result using z01.PrintRune
	printInt(acc)

	// Print a newline character to move to the next line
	z01.PrintRune('\n')
}

// printInt recursively prints the digits of an integer using z01.PrintRune
func printInt(num int) {
	// Handle negative numbers
	if num < 0 {
		z01.PrintRune('-')
		num = -num
	}

	// Recursively print the digits from the most significant to the least significant
	if num >= 10 {
		printInt(num / 10)
	}

	// Print the current digit by converting it to a character
	z01.PrintRune(rune(num%10 + '0'))
}

func main() {
	// Define functions for multiplication, addition, and division
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}

	// Create a slice of integers for testing
	as := []int{500, 2}

	// Test ReduceInt function with multiplication, addition, and division functions
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
