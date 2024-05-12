package main

import "fmt"

func main() {
	// example 1 of fmt.scan
	var name string
	var address int

	fmt.Scan(&name)
	fmt.Scan(&address)

	fmt.Printf("my name is %v and i live in house number %d\n", name, address)
	// output based on the terminals input;
	// inputs; Allan
	// inputs; 30
	// total output on a newline; my name is Allan and i live in house number 30

	//example 2
	/*var name string
	var address int
	var rent float32
	var rating_value bool

	fmt.Scan(&name, &address, &rent, &rating_value)
	fmt.Printf("%s %d %g %t", name, address, rent, rating_value) */
	//inputs; allan
	//40
	//5.4
	// true
	//total output; allan 40 5.4 true
}
