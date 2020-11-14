package listeners

import (
	"time"
)

/* 
InputDelay will stop additional input from happening (to avoid accidental double clicks or hardware issues) 
InputDelay.Block() is the method that will Block for the duration
*/
type InputDelay struct {
	length  time.Duration
	blocked bool
}

// NewInputDelay constructor
func NewInputDelay(milliseconds int) *InputDelay {
	return &InputDelay{length: time.Millisecond * time.Duration(milliseconds), blocked: false}
}

// IsBlocked function check if the input should be blocked
func (d *InputDelay) IsBlocked() bool {
	return d.blocked
}

// Block then unblocks after the delay
func (d *InputDelay) Block() {
	d.blocked = true
	go d.unblockWithDelay()
}

func (d *InputDelay) unblockWithDelay() {
	time.Sleep(d.length)
	d.blocked = false
}
