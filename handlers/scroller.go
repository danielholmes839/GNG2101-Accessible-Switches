package handlers

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
)

// Scroller struct
type Scroller struct {
	delta      int // amount to change the variable value by
	delay      time.Duration
	horizontal bool
	constant   int // the variable that is constant
	variable   int // the variable that wil change. if horizontal is true the x will change
	limit      int // max width when horizontal, max height when verticle
	scrolling  bool
}

// NewScroller constructor
func NewScroller(pixels int, delay time.Duration, reverse bool, horizontal bool) *Scroller {
	var delta int
	if reverse {
		delta = -pixels
	} else {
		delta = pixels
	}

	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	if horizontal {
		return &Scroller{delta: delta, delay: delay, horizontal: horizontal, constant: height / 2, variable: 0, limit: width}
	}
	return &Scroller{delta: delta, delay: delay, horizontal: horizontal, constant: width / 2, variable: 0, limit: height}
}

func (s *Scroller) move() {
	s.variable += s.delta

	// Keep the mouse inside the screen
	if s.variable < 0 {
		s.variable = s.limit
	} else {
		s.variable %= s.limit
	}

	if s.horizontal {
		// Horizontal
		robotgo.MoveMouse(s.variable, s.constant)
	} else {
		// Verticle
		robotgo.MoveMouse(s.constant, s.variable)
	}
}

// Scroll until any input is received on the channel
func (s *Scroller) Scroll(stop <-chan int) {
	scrolling := true
	go func() {
		for scrolling {
			s.move()
			time.Sleep(s.delay)
		}
	}()

	for value := range stop {
		if value == 1 {
			scrolling = false
			break
		}
	}
}

// Reset method
func (s *Scroller) Reset() {
	s.variable = 0
}
