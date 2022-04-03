package main

import (
	"fmt"
	"log"

	"example.com/hello"
)

func main() {
	log.SetPrefix("hello: ")
	log.SetFlags(0)

	names := []string{"Brandon", "Hugo", "Victor"}
	messages, err := hello.HelloAll(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
