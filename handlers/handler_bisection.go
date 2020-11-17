package handlers

import (
	"fmt"

	"github.com/danielholmes839/GNG2101-Switches/handlers/clicker"
	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
)

// Selector struct
type Selector struct {
	min int
	max int

	variable   int
	constant   int
	horizontal bool
}

// NewSelector constructor
func NewSelector(horizontal bool) *Selector {
	selector := &Selector{horizontal: horizontal}
	selector.Reset()
	return selector
}

// Reset method
func (s *Selector) Reset() {
	s.min = 0
	if s.horizontal {
		s.max = int(win.GetSystemMetrics(win.SM_CXSCREEN))
		s.variable = s.max / 2
	} else {
		s.max = int(win.GetSystemMetrics(win.SM_CYSCREEN))
		s.variable = s.max / 2
	}
}

// ChooseMin method
func (s *Selector) ChooseMin() {
	s.max = s.variable
	s.variable = (s.min + s.max) / 2
	s.move()
}

// ChooseMax method
func (s *Selector) ChooseMax() {
	s.min = s.variable
	s.variable = (s.min + s.max) / 2
	s.move()
}

// Move method
func (s *Selector) move() {
	if s.horizontal {
		robotgo.Move(s.variable, s.constant)
	} else {
		robotgo.Move(s.constant, s.variable)
	}
}

// BisectionHandler struct
type BisectionHandler struct {
	state     int
	shortcut  int
	selector  *Selector
	selector1 *Selector
	selector2 *Selector
	clicker   *clicker.Clicker
}

// NewBisectionHandler Constructor
func NewBisectionHandler(shortcut int, horizontalFirst bool) *BisectionHandler {
	handler := &BisectionHandler{
		shortcut:  shortcut,
		selector1: NewSelector(horizontalFirst),
		selector2: NewSelector(!horizontalFirst),
		clicker:   clicker.NewClicker(),
	}
	handler.reset()
	return handler
}

// Handle input
func (h *BisectionHandler) Handle(input <-chan int) {
	for value := range input {
		switch value {
		case 1:
			h.command1()
		case 2:
			h.command2()
		case 3:
			h.command3()
		case 4:
			h.command4()
		}
	}
}

// Command1 func
func (h *BisectionHandler) command1() {
	if h.state == 2 {
		h.reset()
	} else {
		h.selector.ChooseMin()
	}
}

// Command2 func
func (h *BisectionHandler) command2() {
	if h.state == 2 {
		x, y := h.getPos()
		h.clicker.SpecificClick(x, y, h.shortcut)
	} else {
		h.selector.ChooseMax()
	}
}

// Command3 func
func (h *BisectionHandler) command3() {
	if h.state == 2 {
		x, y := h.getPos()
		h.clicker.Click(x, y)
	} else if h.state == 1 {
		h.state = 2
	} else {
		h.state = 1
		h.selector2.constant = h.selector1.variable
		h.selector = h.selector2
	}
}

// Command4 func
func (h *BisectionHandler) command4() {
	h.clicker.Next()
}

// Reset the state
func (h *BisectionHandler) reset() {
	h.selector1.Reset()
	h.selector2.Reset()
	h.selector1.constant = h.selector2.variable
	h.selector = h.selector1
	h.state = 0
}

func (h *BisectionHandler) getPos() (int, int) {
	var x, y int
	if h.selector1.horizontal { // horizontal
		x = h.selector1.variable
		y = h.selector2.variable
	} else {
		x = h.selector2.variable
		y = h.selector1.variable
	}
	fmt.Println(x, y)
	return x, y
}
