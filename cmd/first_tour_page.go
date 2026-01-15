package main

import (
	"fmt"
	"log"
)

// https://go.dev/tour/list
func first_tour_page() {
	fmt.Println("Hello World!")
	fmt.Println(`Hello
	World`) // Using backticks for multi-line string, will already be formatted as is

	// In Go, we have the following types: bool,
	// string,
	// int (and its variations): int8,  int16,  int32,  int64
	// uint (and its variations): uint8, uint16, uint32, uint64, uintptr (which is not recommended for use, just in very specific cases[1])
	// byte (alias for uint8)
	// rune (alias for int32)
	// represents a Unicode code point
	// float32 and float64
	// complex64 and complex128

	// [1]: https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang
	// https://go.dev/tour/basics/11

	var age uint8 = 20
	var height float32 = 1.73
	var name string = "Gabriel"
	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)

	fmt.Printf("Is my age a even number? %t\n", isEven(int(age)))
	fmt.Printf("How old will I be in 20 years from now? %d\n", sum(int(age), 20))

	var ageUser uint8
	fmt.Print("Type up your age: ")
	getAgeInput(&ageUser)
	fmt.Printf("How old will YOU be in 20 years from now? %d\n", sum(int(ageUser), 20))

	var xFH, xSH int = half(2)
	fmt.Printf("If there are 2 oranges, and we share them... I will get %d orange and you will get %d orange.\n", xFH, xSH)

	implicitTypeAndDeclaration := "This string was declared using the ':=' operator, which does not require typing the keyword 'var' nor its type." // Cannot reassign value for this variable if its not an integer
	// Cannot create an constant using ":=" operator
	fmt.Println(implicitTypeAndDeclaration)

	const helloWorld = "Olá Mundo!"
	const olaMundo string = "Hello World!"
	fmt.Println("Invoking value of two constants: " + helloWorld + " " + olaMundo)
}

// func keyword + name of the function + (varName type) + return type + block of code
// https://go.dev/tour/basics/4
func isEven(x int) bool {
	return x%2 == 0
}

// Since x and y share their type, I can ommit x type
// https://go.dev/tour/basics/5
func sum(x, y int) int {
	return x + y
}

func half(x int) (xFirstHalf, xSecondHalf int) {
	xFirstHalf = x / 2
	xSecondHalf = x / 2
	return // Returns xFirstHalf and xSecondHalf using naked return (https://go.dev/tour/basics/7)
}

func getAgeInput(ageUser *uint8) {
	_, err := fmt.Scan(ageUser)
	if err != nil {
		log.Fatalln("Invalid user input! It does need to be an simple unsigned integer within 8 bytes.") // Terminates the program
	}
}

/*
- TODO:

- Continue the tour
https://go.dev/tour/list

- Read about errors in Golang
https://go.dev/doc/faq#exceptions
https://go.dev/blog/error-handling-and-go

- Read more about OOP in Golang, since it does not have classes
*/
