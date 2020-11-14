package main

import (
	//"github.com/go-vgo/robotgo"
	"fmt"

	"github.com/danielholmes839/GNG2101-Switches/handlers"
	"github.com/danielholmes839/GNG2101-Switches/listeners"
)

func main() {
	input := make(chan int)
	listener := listeners.NewSerialListener("COM5", 9600)
	handler := handlers.NewScrollingHandler()

	go listener.Listen(input, 50)
	go handler.Handle(input)

	fmt.Println("Listening...")
	fmt.Scanln()
}
