package greetings

import (
	"errors"
	"fmt"

	"math/rand"
	//"golang.org/x/text/message"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

	/*
		var message string
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
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

func randomFormat() string {
	formats := []string{
		"hello my friend %v",
		"%v its so cool that we've met",
		"buddy, %v, so nice to meet you ",
	}
	return formats[rand.Intn(len(formats))]
}
