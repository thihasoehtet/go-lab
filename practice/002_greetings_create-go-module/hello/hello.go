package main

import (
	"fmt"

	"practice/002_greetings_create-go-module/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
