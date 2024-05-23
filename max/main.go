package main

import "fmt"

// Define a function named Max that takes a slice of integers as input and returns the maximum value.c
func Max(a []int) int {
	// Check if the input slice is empty. If it is, return 0 as there is no maximum value to find.
	if len(a) == 0 {
		return 0
	}
	// Initialize the maximum value as the first element of the slice.
	max := a[0]
	// Iterate over the slice to find the maximum value.
	for _, num := range a {
		// Update the maximum value if a larger value is encountered.
		if num > max {
			max = num
		}
	}
	// Return the maximum value found in the slice.
	return max
}

func main() {
	// Create a slice of integers for testing.
	a := []int{26, 56, 75, 35, 123}
	// Call the Max function to find the maximum value in the slice.
	max := Max(a)
	// Print the maximum value to the console.
	fmt.Println(max)
}
