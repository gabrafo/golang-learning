package main

import (
	"fmt"
)

// https://go.dev/tour/list
func main() {
	// https://go.dev/tour/moretypes/1
	fmt.Println("Pointer manipulation...")

	var q int = 1
	var p *int = &q // Pointer p points to q's memory address.
	// If i hadnt declared that p receives q's address, it would just point to nil as its standard value.
	// I could just start a new integer using the new keyword. Example: var p *int = new(int).
	// Go has an garbage collector, so there is no need to free the memory allocation at the end of the program.

	fmt.Println("'p' points to q, which has ", &q, " as his memory address.")
	fmt.Println("'p' memory address is ", &p, ".")
	fmt.Println("'q' holds the value ", q, ".")
	fmt.Println("I can also access this same value using 'p'. Here it is: ", *p, ".")

	// Structs in Go (OOP)
	// https://go.dev/tour/moretypes/2
}
