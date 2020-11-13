package main

import (
	//"github.com/go-vgo/robotgo"
	"fmt"

	"github.com/danielholmes839/GNG2101-Switches/handlers"
	"github.com/danielholmes839/GNG2101-Switches/helpers"
	"github.com/tarm/serial"
)

func read(input chan int, delay int) {
	d := helpers.NewInputDelay(delay)
	c := &serial.Config{Name: "COM5", Baud: 9600}
	s, _ := serial.OpenPort(c)

	// Read serial input
	buf := make([]byte, 1)
	for {
		s.Read(buf)
		if !d.IsBlocked() {
			// Block input
			d.SetBlocked(true)
			go d.UnblockWithDelay()

			// Receive input
			input <- int(buf[0])
		}
	}
}

func actions(input chan int, handler handlers.Handler) {
	for button := range input {
		switch button {
		case 1:
			handler.Command1()
		case 2:
			handler.Command2()
		case 3:
			handler.Command3()
		case 4:
			handler.Command4()
		}
	}
}

func main() {
	input := make(chan int)
	handler := handlers.NewScrollingHandler() //handlers.EmptyHandler{}
	go handler.Start()
	go read(input, 150)        // send serial input accross the channel with minimum delay between inputs
	go actions(input, handler) // receives serial input and trigger commands of handlers

	fmt.Println("Running!")
	fmt.Scanln()
}
