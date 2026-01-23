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

	// Under the hood, an interface value is a tuple (type, value).
	// The 'type' stores the descriptor of the concrete type and
	// the 'value' holds a copy or pointer to the original data.
	// It is this structure that allows Go to identify and execute the correct method at runtime.
	fmt.Println("Interfaces in Go...")

	// The assignment var g Geometry = Square{10, 10} works because the area and perimeter methods
	// have value receivers (s Square). If one of these methods required a pointer receiver (s *Square),
	// the Square value type would not satisfy the Geometry interface, because a value of type T
	// does not possess the methods that require a *T pointer. In that case, only a pointer
	// (&Square{...}) would satisfy the contract.
	// Our tuple 'g' will hold (Square, {10, 10})
	var g Geometry = Square{10, 10}

	fmt.Println("Structs 'Square' and 'Triangle' implement 'Geometry' interface by having all its methods coded.")
	fmt.Println("Let's see the variable 'g' which type is 'Geometry' and was initialized using a 'Square' literal. g: ", g)
	checkType(g)

	// Dynamic dispatch: At runtime, Go inspects the dynamic type inside the
	// interface to locate the specific method implementation for that concrete
	// type and executes it using the dynamic value as the receiver.
	fmt.Println("Example of dynamic dispatching. Let's see the result of our area function: ", g.area(), ". We used the 'Square' implementation!")

	g = Triangle{10, 3, 4, 5}
	fmt.Println("Let's see the variable 'g' which type is 'Geometry' and has received a 'Triangle' literal. g: ", g)
	checkType(g)
	fmt.Println("Another example of dynamic dispatching. Let's see the result of our area function: ", g.area(), ". We used the 'Triangle' implementation!")
}

// "Square" has public visibility since its first letter is upper-cased
// The struct is exported and can be used in other files in the same package, but not their attributes, which start with an lower-cased letter
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

// Interfaces in Go
// Interfaces define a set of method signatures that a type must implement. They simply act as a contract.
// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
// https://go.dev/tour/methods/9
// https://medium.com/@danielabatibabatunde1/simplifying-structs-methods-and-interfaces-in-golang-e86a0c4618aa
type Geometry interface {
	area() float64
	perimeter() float64
}

// Square already implements area(), now I will make it implement perimeter() as well. That way, it will satisfy the "Geometry" interface contract.
func (s Square) perimeter() float64 {
	return (s.length + s.height) * 2
}

type Triangle struct {
	Height, A, B, C float64 // Attributes are public acessible, just like the struct itself
	// Considering C as basis always
}

// Triangle satisfies Geometry's contract as well
func (t Triangle) area() float64 {
	return (t.C * t.Height) / 2
}

func (t Triangle) perimeter() float64 {
	return t.A + t.B + t.C
}

// Helper function
func checkType(g Geometry) {
	// Checking the first space of our tuple (type)
	fmt.Println("g dynamic type...")

	// Type assertions using a switch
	// https://go.dev/tour/methods/16
	switch g.(type) {
	case Square:
		fmt.Println("g type is 'Square'!")
	case Triangle:
		fmt.Println("g type is 'Triangle'!")
	default:
		fmt.Println("Unknown 'Geometry'.")
	}
}

// TODO
// - Explain concrete and non-concrete types in Go using interfaces as example
