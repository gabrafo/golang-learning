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
		V int // Number of vertexex
		E int // Number of edges
	}
	var nonAnonGraph Graph
	nonAnonGraph.V = 12
	nonAnonGraph.E = 12

	// Interesting Stack Overflow post about constructors in Go (there are no constructors in Go by default): https://stackoverflow.com/questions/18125625/constructors-in-go
	// Another way of declaring variables in Go is listed here: https://go.dev/tour/moretypes/5
	var (
		anotherNonAnonGraph = Graph{12, 12}
	)
	fmt.Println("Example of an object in Go, an graph. Number of vertexes: ", anotherNonAnonGraph.V, ". Number of edges: ", anotherNonAnonGraph.V, ".")

	var a [2]string
	a[0] = "Arrays cannot be resized."
	a[1] = "They have their length declared when they are created."

	// Slicing in Go
	// https://go.dev/doc/effective_go#slices
	// https://go.dev/blog/slices-intro
	// Slices in Go are descriptors (collection of metadata) for arrays that give them more flexibility when it comes to capacity.
	// They hold pointers to an existing or brand new array and have two additional fields: capacity, and length.
	var sliceVariable []string = a[:] // Initializing an slice with natural capacity of 2 (since it is pointing to "a")
	fmt.Println("Full array (accessed via slice): ", sliceVariable)
	sliceVariable[0] = "Modified through slice" // Slices and arrays share memory
	fmt.Println("Array modified: ", a)

	var anotherSliceVariable []string = a[0:1] // Just the first element
	fmt.Println("New first element of the array: ", anotherSliceVariable)
	fmt.Println("Slice additional fields: length ", len(anotherSliceVariable), "; capacity ", cap(anotherSliceVariable), ".")

	var sliceWithBrandNewArray = []string{"A slice is formed by specifying two indices, a low and high bound, separated by a colon: a[low : high]", "This selects a half-open range which includes the first element, but excludes the last one."} // Instantiates an new array with 2 starting positions already filled. Our variable is a slice that points into it. It is an slice literal.
	fmt.Println("Here is an slice literal: ", sliceWithBrandNewArray)

	// make() creates a slice.
	// The slice type and length are mandatory.
	// The capacity is optional; if omitted, it equals the length.
	var sliceWithoutReallocation = make([]int, 3, 5)
	sliceWithoutReallocation[0] = 1
	sliceWithoutReallocation[1] = 2
	sliceWithoutReallocation[2] = 3

	fmt.Println("Slice created with make(): ", sliceWithoutReallocation)
	fmt.Println("Slice length: ", len(sliceWithoutReallocation), "; slice capacity: ", cap(sliceWithoutReallocation), ".")

	// Saving the address of the first element to observe reallocation.
	var beforeAppendPtr = &sliceWithoutReallocation[0]
	fmt.Println("Underlying array address before append: ", beforeAppendPtr)

	// append() receives a slice (mandatory) and one or more values of the same type.
	// It returns a slice, because it may point to a new array.
	// append uses the existing underlying array while capacity allows it.
	// If capacity is exceeded, a new underlying array is allocated.
	// https://go.dev/ref/spec#Appending_and_copying_slices

	// This append fits in the existing capacity.
	sliceWithoutReallocation = append(sliceWithoutReallocation, 4)

	var afterAppendPtr = &sliceWithoutReallocation[0]
	fmt.Println("Slice after append: ", sliceWithoutReallocation)
	fmt.Println("Underlying array address after append: ", afterAppendPtr)
	fmt.Println("Pointer unchanged: ", beforeAppendPtr == afterAppendPtr)

	// Now an array where length and capacity start equal.
	// A slice which points to it will force reallocation on append.
	var experimentalArray = [3]int{1, 2, 3}
	var sliceWithReallocation = experimentalArray[:]

	fmt.Println("New slice literal: ", sliceWithReallocation)
	fmt.Println("Slice length: ", len(sliceWithReallocation), "; slice capacity: ", cap(sliceWithReallocation), ".")

	beforeAppendPtr = &sliceWithReallocation[0]
	fmt.Println("Underlying array address before append: ", beforeAppendPtr)

	// This append exceeds capacity and causes a new array allocation.
	sliceWithReallocation = append(sliceWithReallocation, 4)

	afterAppendPtr = &sliceWithReallocation[0]
	fmt.Println("Slice after append: ", sliceWithReallocation)
	fmt.Println("Underlying array address after append: ", afterAppendPtr)
	fmt.Println("Pointer unchanged: ", beforeAppendPtr == afterAppendPtr)

	// The old array remains unchanged.
	fmt.Println("Original array after an append that exceeds his capacity: ", experimentalArray)
	fmt.Println("After an slice append that exceeds the capacity of the array it is pointing to, the slice creates another array to replace the old one with the changed values, capacity, and length.")

	// Maps
	// https://go.dev/tour/moretypes/19
	var mapInGo map[string]int = make(map[string]int)
	mapInGo["one"] = 1
	fmt.Println("The key 'one' returns ", mapInGo["one"], " in our current map.")

	// Map literal
	// https://go.dev/tour/moretypes/20
	var newMap = map[string]int{
		"one":   11,
		"two":   2,
		"three": 3,
	}
	fmt.Println("The key 'one' returns ", newMap["one"], " in our new map.")

	delete(newMap, "one")

	// I will not go further in map operations, since they are basically the same for every language. If I ever need to use them and don't find quickly how, I will rely on documentation.o
	// Here is more on this subject (aside from the doc itself): https://go.dev/tour/moretypes/22

	// Functions as values
	pow := func(x, y int) (res int) {
		res = 1
		for range y {
			res *= x
		}
		return
	}

	fmt.Println("5 to the power of 5 is: ", pow(5, 5))
}
