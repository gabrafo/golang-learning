package main

import (
	"fmt"
)

// https://go.dev/tour/list
func main() {
	// https://go.dev/tour/methods/1
	fmt.Println("Methods in Go...")

	s := Square{
		height: 12,
		length: 20,
	}
	fmt.Println("Variable 's': ", s)
	fmt.Println("Example of method calling for 'Square' struct: ", s.area())

	var mF MyFloat = 8.5
	fmt.Println("Example of method for non-struct types. I will truncate a number that belongs to 'MyFloat' type (a wrapper of float64). Here we go: ", mF.truncate())

	s.scale(2)
	fmt.Println("I doubled the scale of our 's' variable using a brand new method called 'scale'. It accesses the object via a pointer and changes its attributes without explicit return.")
	fmt.Println("Variable 's': ", s)
}

type Square struct {
	height, length float64
}

// Method for "Square" struct
// Its declaration is different from a regular function
// The parameter "s" is a receiver, basically a "this" keyword in some other languages (Java for instance)
// A method can only have one receiver
// Long story short: we cannot build methods for types not defined in the same package (such as built-in types).
// Because built-in types are not defined by us, we must wrap them, as shown later.
func (s Square) area() float64 {
	return s.height * s.length
}

// I can build a method for non-struct types as well
type MyFloat float64

func (f MyFloat) truncate() MyFloat {
	return MyFloat(int(f))
}

// Pointer as a receiver
// With this, I can directly change the value of the struct fields
// It may also avoid copying the value, especially for large structs.
// https://go.dev/tour/methods/4
// https://go.dev/tour/methods/8
// About calling methods that have a pointer as a receiver: https://go.dev/tour/methods/6
func (s *Square) scale(x float64) {
	s.height *= x
	s.length *= x
}
