package main

import (
	"fmt"
)

// https://go.dev/tour/list
func main() {
	// For in Go (https://go.dev/tour/flowcontrol/1)
	fmt.Println("Counting to three...")
	for i := 0; i < 3; i++ {
		fmt.Println(i+1, "...")
	}

	// Another (more idiomatic) way of doing the exact same thing
	for i := range 3 {
		fmt.Println(i+1, "...")
	}
}
