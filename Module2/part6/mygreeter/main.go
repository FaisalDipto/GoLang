package main

import (
	"log"

	"mygreeter/greetings"

	"github.com/fatih/color"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message with a valid name.
	message, err := greetings.Hello("Faisal")
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message in green.
	color.Green(message)
}