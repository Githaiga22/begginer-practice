package main

import "fmt"

func main() {
	name := "John"
	age := 30

	// fmt.Printf() prints the formatted string to the console
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", name, age)

	// fmt.Sprintf() returns the formatted string
	formattedString := fmt.Sprintf("Hello, my name is %s and I am %d years old.\n", name, age)

	// We can use the formatted string however we want, such as printing it later
	fmt.Println("Formatted string:", formattedString)
}
