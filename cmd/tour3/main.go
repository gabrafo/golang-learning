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

	// Structs in Go (OOP).
	// https://go.dev/tour/moretypes/2
	// https://go.dev/ref/spec#Struct_types
	// Struct is a type in Go. A struct is defined as an composed value that groups up named fields, which of them with its particular types.
	// An struct cannot hold on another struct.
	var graph struct {
		X int
		Y int
	}
	graph.Y = 12
	graph.X = 12
	// This is an anon struct. Basically an object with no class. There is no easy way to recreate its structure in other objects.

	// Here is how you can make an struct reusable by other objects.
	type Graph struct {
		X int
		Y int
	}
	var nonAnonGraph Graph
	nonAnonGraph.X = 12
	nonAnonGraph.Y = 12

	// Interesting Stack Overflow post about constructors in Go (there are no constructors in Go by default): https://stackoverflow.com/questions/18125625/constructors-in-go
	// Another way of declaring variables in Go is listed here: https://go.dev/tour/moretypes/5
	var (
		anotherNonAnonGraph = Graph{12, 12}
	)

}
