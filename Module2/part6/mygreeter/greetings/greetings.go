package greetings

import (
	"errors" // Make sure to add this!
	"fmt"
)

// Hello now returns a string AND an error.
func Hello(name string) (string, error) {
	// Use the private function to check for a valid name.
	if !isValidName(name) {
		// If the name is invalid, return an empty string and an error.
		return "", errors.New("empty name is not allowed")
	}

	// If the name is valid, return a greeting and a 'nil' error.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

// (Your private isValidName function stays the same)
func isValidName(name string) bool {
	return name != ""
}