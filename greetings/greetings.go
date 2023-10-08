package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	greetings := make(map[string]string)

	for _, name := range names {
		greeting, err := Hello(name)

		if err != nil {
			return nil, err
		}

		greetings[name] = greeting
	}

	return greetings, nil
}

func randomFormat() string {
	var formats = []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Bonjour, %v!",
		"Hola, %v!",
		"Bom dia, %v!",
	}

	return formats[rand.Intn(len(formats))]
}
