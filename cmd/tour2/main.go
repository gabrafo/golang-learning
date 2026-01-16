package main

import (
	"fmt"
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
}
