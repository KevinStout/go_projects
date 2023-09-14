package main

import (
	"fmt"

	"b.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Kevin")
	fmt.Println(message)
}