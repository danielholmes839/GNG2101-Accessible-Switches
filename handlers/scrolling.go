package handlers

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
)

// Scroller struct
type Scroller struct {
	speed     int
	direction int // 0 -> horizontal, 1 -> vertical
	constant  int // the variable that is constant
	variable  int // the variable that wil change. if horizontal is true the x will change
	limit     int // max width when horizontal, max height when verticle
}

func (s *Scroller) scroll() {
	s.variable += s.speed
	s.variable %= s.limit
	if s.direction == 0 {
		robotgo.MoveMouse(s.variable, s.constant)
	} else {
		robotgo.MoveMouse(s.constant, s.variable)
	}
}

// ScrollingHandler struct
type ScrollingHandler struct {
	running   bool
	scroller1 *Scroller
	scroller2 *Scroller
}

// NewScrollingHandler Constructor
func NewScrollingHandler() *ScrollingHandler {
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))
	return &ScrollingHandler{
		running:   false,
		scroller1: &Scroller{speed: 20, direction: 0, constant: height / 2, variable: 0, limit: width},
		scroller2: &Scroller{speed: 2, direction: 1, constant: width / 2, variable: 0, limit: height},
	}
}

// Command1 func (stop)
func (handler *ScrollingHandler) Command1() {
	handler.running = false
}

// Command2 func (click)
func (handler ScrollingHandler) Command2() {
}

// Command3 func
func (handler ScrollingHandler) Command3() {
}

// Command4 func
func (handler ScrollingHandler) Command4() {
}

// Start func
func (handler *ScrollingHandler) Start() {
	handler.running = true
	for handler.running {
		handler.scroller1.scroll()
		time.Sleep(time.Millisecond * 15)
	}

	handler.scroller2.constant = handler.scroller1.variable
	handler.running = true
	for handler.running {
		handler.scroller2.scroll()
		time.Sleep(time.Millisecond * 15)
	}
}
