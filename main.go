package main

import (
	"fmt"
	"github.com/danielholmes839/GNG2101-Switches/handlers"
	"github.com/danielholmes839/GNG2101-Switches/listeners"
)

func main() {
	config := handlers.GetConfig()
	listener := listeners.NewSerialListener(config.SerialPort, 9600)
	handler, _ := handlers.GetHandler(config)
	
	input := make(chan int)
	go listener.Listen(input, config.InputDelay)
	fmt.Println("Listening...")
	
	handler.Handle(input)

	
}
