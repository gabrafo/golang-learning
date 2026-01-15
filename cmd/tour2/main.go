package main

import (
	"fmt"
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
		fmt.Println("In Go there are no needs for parathensis in if conditions, but brackets are essential.")
	}

	// https://go.dev/tour/flowcontrol/6
	if j := 2; j < i {
		fmt.Println("Just like in for, the if statement can start with a short statement to execute before the condition. Any variables declared within this statement are local, and just exist in the if scope.")
	}
}
