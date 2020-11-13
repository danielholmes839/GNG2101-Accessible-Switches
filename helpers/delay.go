package helpers

import (
	"time"
)

// InputDelay will stop additional input from happening (to avoid accidental double clicks or hardware issues)
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

// SetBlocked function check if the input should be blocked
func (d *InputDelay) SetBlocked(value bool) {
	d.blocked = value
}

// UnblockWithDelay lock function start blocking inputs
func (d *InputDelay) UnblockWithDelay() {
	time.Sleep(d.length)
	d.SetBlocked(false)
}
