package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	var name string
	flag.StringVar(&name, "name", "world", "The name to greet")

	// Parse flags
	flag.Parse()

	// Greet the user
	fmt.Println("Hello,", name)
}
