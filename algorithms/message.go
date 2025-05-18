package algorithms

import (
	"errors"
	"fmt"
	"math/rand"
)

// public beacuse it starts with an upper case letter
func Message() (string, error) {
	message := randomFormat()
	if message == "" {
		return "", errors.New("mensagem vazia")
	}

	return message, nil
}

// private because it starts with a lower case letter
func randomFormat() string {
	formats := []string{
		"Ola Bem vindo",
		"Bom ver voce.",
		"Prazer em te conhecer",
		"",
		"",
	}
	return formats[rand.Intn(len(formats))]
}

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
