package handlers

import (
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
	if s.horizontal {
		s.max = s.variable
		s.variable = (s.min + s.max) / 2
		s.move()
	} else {
		s.ChooseMax()
	}
}

// ChooseMax method
func (s *Selector) ChooseMax() {
	if s.horizontal {
		s.min = s.variable
		s.variable = (s.min + s.max) / 2
		s.move()
	} else {
		s.ChooseMin()
	}
}

// Move method
func (s *Selector) move() {
	if s.horizontal {
		robotgo.Move(s.variable, s.constant)
	} else {
		robotgo.Move(s.constant, s.variable)
	}
}
