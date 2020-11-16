package handlers

import (
	"fmt"

	"github.com/lxn/win"
)

// Selector struct
type Selector struct {
	mid int
	min int
	max int

	constant   int
	variable   int
	horizontal bool
}

// NewSelector constructor
func NewSelector(horizontal bool) *Selector {
	selector := &Selector{}
	selector.Reset()
	return selector
}

// Reset method
func (s *Selector) Reset() {
	s.min = 0
	if s.horizontal {
		s.max = int(win.GetSystemMetrics(win.SM_CXSCREEN))
		s.mid = s.max / 2
	} else {
		s.max = int(win.GetSystemMetrics(win.SM_CYSCREEN))
		s.mid = s.max / 2
	}
}

// ChooseMin method
func (s *Selector) ChooseMin() {
	s.max = s.mid
	s.mid = (s.min + s.max) / 2
}

// ChooseMax method
func (s *Selector) ChooseMax() {
	s.min = s.mid
	s.mid = (s.min + s.max) / 2
}

// BisectionHandler struct
type BisectionHandler struct {
	onFirstSelection bool
	selector1        *Selector
	selector2        *Selector
	command          chan int
}

// NewBisectionHandler Constructor
func NewBisectionHandler() *BisectionHandler {
	return &BisectionHandler{}
}

// Handle input
func (h *BisectionHandler) Handle(input chan int) {
	defaultHandle(input, h)
}

// Command1 func (stop)
func (h *BisectionHandler) Command1() {
	if h.onFirstSelection {
		h.selector1.ChooseMin()
	} else {
		h.selector2.ChooseMin()
	}
}

// Command2 func (click)
func (h BisectionHandler) Command2() {
	if h.onFirstSelection {
		h.selector2.ChooseMin()
	} else {
		h.selector2.ChooseMin()
	}
}

// Command3 func
func (h BisectionHandler) Command3() {
	if h.onFirstSelection {
		fmt.Println("Second selection")
		h.onFirstSelection = false
	} else {
		h.onFirstSelection = true
		// some click operation
	}
}

// Command4 func
func (h BisectionHandler) Command4() {
}

// Start func
func (h *BisectionHandler) start() {
}
