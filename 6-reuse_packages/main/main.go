package main

import (
	"fmt"

	"example.com/hello"
)

func main() {
	message := hello.Hello("Brandon")
	fmt.Println(message)
}
