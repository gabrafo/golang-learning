package main

import (
	"fmt"
	"os"
	"runtime"
)

// https://go.dev/tour/list
func main() {
	// For in Go
	// https://go.dev/tour/flowcontrol/1
	// https://go.dev/tour/flowcontrol/2
	fmt.Println("Counting to three...")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d...\n", i+1)
	}

	fmt.Println("---")

	// Another (more idiomatic) way of doing the exact same thing
	for i := range 3 {
		fmt.Printf("%d...\n", i+1)
	}

	fmt.Println("---")

	// Yet another way
	// There is no While in Go, just For
	// https://go.dev/tour/flowcontrol/3
	var i int = 0
	for i < 3 {
		fmt.Printf("%d...\n", i+1)
		i++
	}

	// https://go.dev/tour/flowcontrol/5
	condition := true
	if condition {
		fmt.Println("In Go there is no need for parathensis in if conditions, but brackets are essential.")
	}

	// https://go.dev/tour/flowcontrol/6
	if j := 2; j < i {
		fmt.Println("Just like in for, the if statement can start with a short statement to execute before the condition. Any variables declared within this statement are local, and just exist in the if/else scope.")
	} else {
		fmt.Printf("In this code, there was an if condition, where it checked if %d is lesser than %d. The condition ended up being truth, that is why this message is not being printed out to you, dear reader. If it is being printed though, maybe you, or me, have changed the values.\n", j, i) // https://go.dev/tour/flowcontrol/7
	}

	// No need to add the "break" clause
	// Evaluates from top to bottom
	// Switch with no condition is just "true"
	// https://go.dev/tour/flowcontrol/9
	// https://go.dev/tour/flowcontrol/10
	// https://go.dev/tour/flowcontrol/11
	fmt.Println("Currently running this code in: ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("An actual operating system!")
	default:
		fmt.Println("Nothing to brag about...")
	}

	fmt.Println(fileHandling())

	fmt.Println("Counting once again...")
	deferExample()
}

func fileHandling() (message string) {
	file, err := os.Create("defer_explanation.txt")

	var pointerToMessage *string = &message

	defer errorHandling(err, pointerToMessage) // The deferred call will run when fileHandling() returns. The arguments are evaluated immediately, but execution is delayed.

	defer file.Close() // Even though deferred statements are executed in a LIFO way, their arguments are evaluated immediately. As such, I cannot place file.Close() before the errorHandling() call. The reason being: if os.Create() fails, file will be a nil variable. If the runtime evaluate an statement "nil.Close()" it will be trouble.

	file.WriteString(`A defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.`)

	return
}

func errorHandling(err error, errMessage *string) {
	if err != nil {
		*errMessage = fmt.Errorf("ERROR: %w\n", err).Error()
	} else {
		*errMessage = "No errors in file handling. New file created."
	}
}

// Naked returns/Regular returns and defer interactions explained in the two links bellow

// https://stackoverflow.com/questions/37248898/how-does-defer-and-named-return-value-work

// https://go.dev/ref/spec#Defer_statements
// This one, specially, explains why fileHandling worked and why I had to use an naked return, and not an regular return.
// "If the surrounding function returns through an explicit return statement, deferred functions are executed after any result parameters are set by that return statement but before the function returns to its caller", that is why an regular return with the "message" variable would not work. Besides, "If the deferred function has any return values, they are discarded when the function completes".
// And here is why I had to use an naked return: "if the deferred function is a function literal and the surrounding function has named result parameters that are in scope within the literal, the deferred function may access and modify the result parameters before they are returned".

// More examples of defer
func deferExample() {
	fmt.Println("0")
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")

	// First prints zero, and then prints in LIFO order, so: 0, 1, 2, 3.
}
