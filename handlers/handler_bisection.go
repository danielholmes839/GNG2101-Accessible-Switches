package handlers

import (
	"fmt"

	"github.com/danielholmes839/GNG2101-Switches/handlers/clicker"
)

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
func NewBisectionHandler(config *BisectionConfig) *BisectionHandler {
	handler := &BisectionHandler{
		selector1: NewSelector(config.HorizontalFirst),
		selector2: NewSelector(config.HorizontalFirst),
		clicker:   clicker.NewClicker(config.Shortcut),
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
		h.clicker.SpecificClick(x, y, h.clicker.Shortcut)
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
