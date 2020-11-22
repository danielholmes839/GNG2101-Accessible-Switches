package handlers

import (
	"time"

	"github.com/danielholmes839/GNG2101-Switches/handlers/clicker"
)

// ScrollingHandler struct
type ScrollingHandler struct {
	clicker   *clicker.Clicker
	scroller1 *Scroller
	scroller2 *Scroller
	command   chan int
}

// NewScrollingHandler Constructor
func NewScrollingHandler(config *ScrollingConfig) *ScrollingHandler {
	delay := time.Duration((1000 / config.FramesPerSecond) * 1000000)
	pixels := config.PixelsPerFrame

	var scroller1Reverse, scroller2Reverse bool
	if config.HorizontalFirst {
		scroller1Reverse = !config.LeftToRight
	} else {
		scroller2Reverse = !config.TopToBottom
	}

	return &ScrollingHandler{
		clicker:   clicker.NewClicker(config.Shortcut),
		scroller1: NewScroller(pixels, delay, scroller1Reverse, config.HorizontalFirst),
		scroller2: NewScroller(pixels, delay, scroller2Reverse, !config.HorizontalFirst),
		command:   make(chan int),
	}
}

// Handle input
func (h *ScrollingHandler) Handle(input <-chan int) {
	go h.start()
	for value := range input {
		h.command <- value
	}
}

// Start func
func (h *ScrollingHandler) start() {
	for {
		h.scroller1.Scroll(h.command)
		h.scroller2.constant = h.scroller1.variable
		h.scroller2.Scroll(h.command)

		x, y := h.getPos()

		for value := range h.command {
			exit := false
			switch value {
			case 1:
				exit = true
			case 2:
				// shortcut click
				h.clicker.SpecificClick(x, y, h.clicker.Shortcut)
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

func (h *ScrollingHandler) getPos() (int, int) {
	var x, y int
	if h.scroller1.horizontal { // horizontal
		x = h.scroller1.variable
		y = h.scroller2.variable
	} else {
		x = h.scroller2.variable
		y = h.scroller1.variable
	}
	return x, y
}
