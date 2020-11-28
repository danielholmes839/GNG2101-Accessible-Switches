package handlers

import (
	"time"

	"github.com/danielholmes839/GNG2101-Switches/handlers/clicker"
)

// ScrollingHandler struct
type ScrollingHandler struct {
	shortcut1 int
	shortcut2 int
	reset     bool
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
		shortcut1: config.Shortcut1,
		shortcut2: config.Shortcut2,
		reset:     config.ScrollPositionResets,
		clicker:   clicker.NewClicker(),
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
				h.clicker.SpecificClick(x, y, h.shortcut1)
			case 3:
				// shortcut click
				h.clicker.SpecificClick(x, y, h.shortcut2)
			case 4:
				// selected click
				h.clicker.Click(x, y)
			case 5:
				// change selected click
				h.clicker.Next()
			}

			if exit {
				break
			}
		}

		if h.reset {
			h.scroller1.Reset()
			h.scroller2.Reset()
		} else {
			h.scroller1.constant = h.scroller2.variable
		}
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
