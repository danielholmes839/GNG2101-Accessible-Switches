package main

import (
	//"github.com/go-vgo/robotgo"
	"github.com/danielholmes839/GNG2101-Switches/handlers"

	"github.com/tarm/serial"
)

func read(input chan int) {
	c := &serial.Config{Name: "COM5", Baud: 9600}
	s, _ := serial.OpenPort(c)

	// Read serial input
	buf := make([]byte, 1)
	for {
		s.Read(buf)
		input <- int(buf[0])
	}

}

func actions(input chan int, handler handlers.Handler) {
	for value := range input {
		switch value {
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
	handler := handlers.DefaultHandler{}
	input := make(chan int)
	go read(input)
	actions(input, handler)

	// robotgo.ScrollMouse(10, "up")
	// robotgo.MouseClick("left", true)
	// robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
}
