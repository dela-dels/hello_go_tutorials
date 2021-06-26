package main

import (
	"deladels/greetings"
	"fmt"
	"log"
)

func main() {

	log.SetPrefix("error encountered:")
	log.SetFlags(0)

	names := []string{"Dela Dels", "Foobah", "Nyams"}

	message, err := greetings.GreetAll(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
