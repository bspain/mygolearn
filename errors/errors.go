package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {  // The error interface { Error() string } is looked for when printing values
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{	// *MyError supports the error interface now
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// A nil error denotes success; a non-nil error denotes failure
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
