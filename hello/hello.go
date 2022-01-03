package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Faizan")
	fmt.Println(message)
}
