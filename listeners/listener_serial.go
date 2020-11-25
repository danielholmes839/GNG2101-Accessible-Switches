package listeners

import (
	"fmt"

	"github.com/tarm/serial"
)

// SerialListener struct
type SerialListener struct {
	port     string
	baudrate int
}

// NewSerialListener constructor
func NewSerialListener(port string, baudrate int) *SerialListener {
	return &SerialListener{port: port, baudrate: baudrate}
}

// Listen for serial input and send it on the channel. with an input delay
func (l *SerialListener) Listen(input chan<- int, delay int) {
	d := NewInputDelay(delay)
	c := &serial.Config{Name: l.port, Baud: l.baudrate}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic("Could not open serial port on port: " + l.port)
	}
	fmt.Println("Opened Serial Port Successfully")
	// Read serial input
	buf := make([]byte, 1)
	for {
		s.Read(buf)
		if !d.IsBlocked() {
			d.Block()
			value := int(buf[0])
			fmt.Println(value)
			input <- value
		}
	}
}
