package listeners

import (
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
	c := &serial.Config{Name: "COM5", Baud: 9600}
	s, _ := serial.OpenPort(c)

	// Read serial input
	buf := make([]byte, 1)
	for {
		s.Read(buf)
		if !d.IsBlocked() {
			d.Block()
			input <- int(buf[0])
		}
	}
}
