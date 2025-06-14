// Package main is the entry point for Go applications
// Every Go application must have a main package
package main

import (
	"fmt"  // fmt package provides functions for formatted I/O operations
	"os"   // os package provides platform-independent interface to operating system functionality
	"time" // time package provides functionality for measuring and displaying time
)

// main is the entry point of the application.
// It simulates a coin toss and prints either "head" or "tail" to the console.
// If 't' parameter is provided, it displays the current time instead.
func main() {
	// Check if command-line arguments are provided
	if len(os.Args) > 1 && os.Args[1] == "t" {
		// Display current time if 't' parameter is provided
		currentTime := time.Now()
		fmt.Println("Current time:", currentTime.Format(time.RFC1123))
		return
	}
}
