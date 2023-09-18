package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a random greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message. or an error if name is empty
	if name == "" {
		return "", errors.New("empty name") 
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns on of a set of greeting messages. The returned message is selected at random. 
func randomFormat() string {
	// a slice of message formats. A slie is like an array, except that its size is dynamically resizable
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
