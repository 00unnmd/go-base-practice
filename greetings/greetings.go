package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("error! Name is empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	
	return message, nil
}

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

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hello, %v. Welcome and Go!",
		"Great to see you, %v! Let's Go",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}