package handlers

import (
	"fmt"
	"time"

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
	} else {
		return &Scroller{delta: delta, direction: 1, constant: width / 2, variable: 5, limit: height}
	}
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

// Scroll until 1 is received on the channel
func (s *Scroller) Scroll(command chan int) {
	running := true
	go func() {
		for running {
			s.move()
			time.Sleep(time.Millisecond * 5)
		}
	}()
	for value := range command {
		if value == 1 {
			running = false
			return
		}
	}
}

func (s *Scroller) reset() {
	s.variable = 0
}

// ScrollingHandler struct
type ScrollingHandler struct {
	scroller1 *Scroller
	scroller2 *Scroller
	command   chan int
}

// NewScrollingHandler Constructor
func NewScrollingHandler() *ScrollingHandler {
	return &ScrollingHandler{
		scroller1: NewScroller(5, false, true),
		scroller2: NewScroller(5, true, false),
		command:   make(chan int),
	}
}

// Handle input
func (h *ScrollingHandler) Handle(input chan int) {
	go defaultHandle(input, h)
	h.start()
}

// Command1 func (stop)
func (h *ScrollingHandler) Command1() {
	h.command <- 1
}

// Command2 func (click)
func (h ScrollingHandler) Command2() {
	h.command <- 2
}

// Command3 func
func (h ScrollingHandler) Command3() {
	h.command <- 3
}

// Command4 func
func (h ScrollingHandler) Command4() {
	h.command <- 4
}

// Start func
func (h *ScrollingHandler) start() {
	for {
		h.scroller1.Scroll(h.command)
		h.scroller2.constant = h.scroller1.variable
		h.scroller1.reset()
		h.scroller2.Scroll(h.command)
		h.scroller2.reset()

		value := <-h.command
		switch value {
		case 2:
			fmt.Println(value)
		case 3:
			fmt.Println(value)
		case 4:
			fmt.Println(value)
		}
	}
}
