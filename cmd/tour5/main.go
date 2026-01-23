package main

import (
	"fmt"
)

// https://go.dev/tour/list
func main() {
	// https://go.dev/tour/generics/1
	fmt.Println("Generics in Go...")

	Index(5, 5)
	Index("Hello", "Hello")
	Index(10, 20)

	var listStrings List[string]
	listStrings.val = "Hello"
	listStrings.next = &List[string]{nil, "World!"}
	listStrings.Explore()

	var listInt List[int]
	listInt.val = 1
	listInt.next = &List[int]{next: &List[int]{nil, 3}, val: 2}
	listInt.Explore()
}

// Go functions can be written to work on multiple types using type parameters.
// The type parameters of a function appear between brackets, before the function's arguments.
// `comparable` is a useful constraint that allows the use of `==` and `!=` operators.
func Index[T comparable](x T, y T) {
	fmt.Printf("Comparing values: %v and %v (type: %T)\n", x, y, x)

	if x == y {
		fmt.Println("Values are equal!")
	} else {
		fmt.Println("Values are different!")
	}
}

// Generic type
// https://go.dev/tour/generics/2
// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Explore() {
	fmt.Println("Exploring a list with generic type using the 'Explore' function...")
	fmt.Printf("Current type: %T\n", l)
	for l != nil {
		fmt.Println(l.val)
		l = l.next
	}
}

