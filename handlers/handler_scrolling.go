package handlers

import (
	"time"

	"github.com/danielholmes839/GNG2101-Switches/handlers/clicker"
	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
)

// Scroller struct
type Scroller struct {
	delta     int // amount to change the variable value by
	direction int // 0 -> horizontal, 1 -> vertical
	constant  int // the variable that is constant
	variable  int // the variable that wil change. if horizontal is true the x will change
	limit     int // max width when horizontal, max height when verticle
	scrolling bool
}

// NewScroller constructor
func NewScroller(speed int, reverse bool, horizontal bool) *Scroller {
	var delta int
	if reverse {
		delta = -speed
	} else {
		delta = speed
	}

	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	if horizontal {
		return &Scroller{delta: delta, direction: 0, constant: height / 2, variable: 5, limit: width}
	}
	return &Scroller{delta: delta, direction: 1, constant: width / 2, variable: 5, limit: height}
}

func (s *Scroller) move() {
	s.variable += s.delta

	// Keep the mouse inside the screen
	if s.variable < 0 {
		s.variable = s.limit
	} else {
		s.variable %= s.limit
	}

	if s.direction == 0 {
		// Horizontal
		robotgo.MoveMouse(s.variable, s.constant)
	} else {
		// Verticle
		robotgo.MoveMouse(s.constant, s.variable)
	}
}

// Scroll until any input is received on the channel
func (s *Scroller) Scroll(stop chan int) {
	scrolling := true
	go func() {
		for scrolling {
			s.move()
			time.Sleep(time.Millisecond * 15)
		}
	}()

	<-stop
	scrolling = false
}

// Reset method
func (s *Scroller) Reset() {
	s.variable = 0
}

// ScrollingHandler struct
type ScrollingHandler struct {
	clicker   *clicker.Clicker
	scroller1 *Scroller
	scroller2 *Scroller
	command   chan int
}

// NewScrollingHandler Constructor
func NewScrollingHandler() *ScrollingHandler {
	return &ScrollingHandler{
		clicker:   clicker.NewClicker(),
		scroller1: NewScroller(5, false, true),
		scroller2: NewScroller(5, true, false),
		command:   make(chan int),
	}
}

// Handle input
func (h *ScrollingHandler) Handle(input chan int) {
	go h.start()
	defaultHandle(input, h)
}

// Command1 func (stop)
func (h *ScrollingHandler) Command1() {
	h.command <- 1
}

// Command2 func (click)
func (h *ScrollingHandler) Command2() {
	h.command <- 2
}

// Command3 func
func (h *ScrollingHandler) Command3() {
	h.command <- 3
}

// Command4 func
func (h *ScrollingHandler) Command4() {
	h.command <- 4
}

// Start func
func (h *ScrollingHandler) start() {
	for {
		h.scroller1.Scroll(h.command)
		h.scroller2.constant = h.scroller1.variable
		h.scroller2.Scroll(h.command)

		var x, y int
		if h.scroller1.direction == 0 { // horizontal
			x = h.scroller1.variable
			y = h.scroller2.variable
		} else {
			x = h.scroller2.variable
			y = h.scroller1.variable
		}

		for value := range h.command {
			exit := false
			switch value {
			case 1:
				exit = true
			case 2:
				// shortcut click
				h.clicker.SpecificClick(x, y, 2)
			case 3:
				// selected click
				h.clicker.Click(x, y)
			case 4:
				// change selected click
				h.clicker.Next()
			}

			if exit {
				break
			}
		}

		h.scroller1.Reset()
		h.scroller2.Reset()
	}
}
