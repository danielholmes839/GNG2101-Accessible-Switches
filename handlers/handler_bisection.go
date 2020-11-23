package handlers

import (
	"fmt"

	"github.com/danielholmes839/GNG2101-Switches/handlers/clicker"
)

// BisectionHandler struct
type BisectionHandler struct {
	state     int // 0 selecting for first orientation, 1 selecting for second orientation, 2 ready to click
	shortcut1 int
	shortcut2 int
	selector  *Selector
	selector1 *Selector
	selector2 *Selector
	clicker   *clicker.Clicker
}

// NewBisectionHandler Constructor
func NewBisectionHandler(config *BisectionConfig) *BisectionHandler {
	handler := &BisectionHandler{
		shortcut1: config.Shortcut1,
		shortcut2: config.Shortcut2,
		selector1: NewSelector(config.HorizontalFirst),
		selector2: NewSelector(!config.HorizontalFirst),
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
		case 5:
			h.command5()
		}
	}
}

// Command1 func
func (h *BisectionHandler) command1() {
	// Reset when ready to click or choose min (left, top)
	if h.state == 2 {
		h.reset()
	} else {
		h.selector.ChooseMin()
	}
}

func (h *BisectionHandler) command2() {
	// Shortcut when ready to click or choose max (right, bottom)
	if h.state == 2 {
		x, y := h.getPos()
		h.clicker.SpecificClick(x, y, h.shortcut1)
	} else {
		h.selector.ChooseMax()
	}
}

func (h *BisectionHandler) command3() {
	// Shortcut when ready to click
	if h.state == 2 {
		x, y := h.getPos()
		h.clicker.SpecificClick(x, y, h.shortcut2)
	}
}

func (h *BisectionHandler) command4() {
	// Toggle state to ready to click
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

func (h *BisectionHandler) command5() {
	// Go to next type of click
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
