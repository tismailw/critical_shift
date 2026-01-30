package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("Error Log: ")
	log.SetFlags(0) //disables extra log info

	names := []string{
		"Tahir",
		"Iqbal",
		"Soha",
	}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err) //exits the loop without returning
	}

	fmt.Println(messages)
}
