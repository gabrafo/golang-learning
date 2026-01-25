package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"os"
	"os/signal"
	"syscall"
)

// Example below based on an exercise from Chapter 8 of "The Go Programming Language"
// https://go.dev/tour/list
func main() {
	// https://go.dev/tour/concurrency/1
	fmt.Println("Concurrency in Go...")

	// Go Routines
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("Failed to start TCP listener:", err) // Ends program, there is no reason for continuing if listening failed
	}

	// Channel that listens to OS signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM) // Notifies "done" that an interrupt signal was sent (ending the server via terminal)

	// Go routine that ends the server, listens to "done" waiting for an interrupt signal
	go endServer(done, listener)

	serve(listener, done)
}

// Blocking function, will execute this loop endlessly till done sends a signal
func serve(listener net.Listener, done chan os.Signal) {
	for { // Endless loop
		conn, err := listener.Accept()
		if err != nil {
			select {
				case <-done:
					log.Println("Listener closed, connections loop exiting.")
            		return 

				// Default case basically is an alternative if none of the selected cases happen.
				// The idea is to make our select non-blocking, since it is not just waiting for signals, and has an alternative to keep on our for loop.
				// https://go.dev/tour/concurrency/6
				default: 
					log.Println("Error accepting connection! Error message: ", err)
					continue // Moves to the next iteration, without executing what comes next in this iteration
			}
		}
		log.Println("Connection accepted!")
		go handleConn(conn, done) // Handling connections concurrently with a Go Routine
	}
}

func handleConn(conn net.Conn, done chan os.Signal) {
	defer conn.Close() // Closes the connection at the end of the function execution
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Println("Client disconnected. Error message: ", err)
			return // Ending Go Routine
		}

		// Select
		// Is similar to a "switch", but each case specifies a communication operation (send or receive) on a channel.
		// Without a "default" case, select blocks the goroutine until one of the cases is ready to proceed.
		// https://go.dev/tour/concurrency/5
		select {

		case <-done:
			// Triggered if a value is received from "done" (or if it is closed), ends active connections.
			log.Println("Stopping handler via signal.")
			return

		case <- ticker.C:
			// Channel that sends a signal after 1 second.
			// This acts as a timeout or wait and just moves to the next execution of our loop. 
		}
	}
}

func endServer(done chan os.Signal, listener net.Listener) {
	<-done // Go routine blocked, waiting signal
	log.Println("Shutting down server...")
	listener.Close() // Freeing 8000
}

// Interesting resources (I plan to deep dive them later on and bring more examples and thoughts to the table):
// - https://stackoverflow.com/questions/48638663/what-is-relationship-between-goroutine-and-thread-in-kernel-and-user-state
// - https://www.youtube.com/watch?v=KBZlN0izeiY&t=536s
// - https://www.reddit.com/r/golang/comments/117a4x7/how_can_goroutines_be_more_scalable_than_kernel/