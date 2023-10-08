package main

import (
	"fmt"
	"log"

	"github.com/tomasrodathg/learning-go/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Get a greeting message and print it.
	names := []string{"Tomas", "Joe", "Melissa", "Mish"}

	res, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
