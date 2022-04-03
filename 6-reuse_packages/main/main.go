package main

import (
	"fmt"
	"log"

	"example.com/hello"
)

func main() {
	log.SetPrefix("hello: ")
	log.SetFlags(0)

	message, err := hello.Hello("Brandon")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
