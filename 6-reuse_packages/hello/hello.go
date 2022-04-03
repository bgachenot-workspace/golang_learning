package hello

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hello, %v!",
		"I'm glad to see u back, %v!",
		"Where were you? There is 3 spiderman now!! %v",
	}

	return formats[rand.Intn(len(formats))]
}
