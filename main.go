package main

import (
	"fmt"

	"github.com/danielholmes839/GNG2101-Switches/handlers"
	"github.com/danielholmes839/GNG2101-Switches/listeners"
)

func main() {
	input := make(chan int)
	listener := listeners.NewSerialListener("COM5", 9600)
	handler := handlers.NewScrollingHandler() 						//handlers.NewBisectionHandler(4, true)

	go listener.Listen(input, 50)
	go handler.Handle(input)

	fmt.Println("Listening...")
	fmt.Scanln()
}
