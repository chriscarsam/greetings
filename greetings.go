package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the specified person.
func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("Empty name")
	}
	// Returns a greeting that includes the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// randomFormat returns one of several greeting message formats
// is randomly selected
func randomFormat() string {
	formats := []string{
		"¡Hello, %v! ¡Welcome!",
		"¡Good to see you, %v!",
		"¡Greetings, %v! ¡Nice to meet you!",
		"¡Good By, %v! ¡See you later!",
	}

	return formats[rand.Intn(len(formats))]
}
