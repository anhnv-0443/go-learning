package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Viet Anh", 30)
	fmt.Println(message)
}
